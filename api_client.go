package magento2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"gopkg.in/resty.v1"
)

const (
	RetryAttempts       = 3
	RetryWaitSeconds    = 5
	RetryMaxWaitSeconds = 20
)

type Client struct {
	HTTPClient *resty.Client
}

// tokenRefresher holds everything needed to transparently refresh an auth token.
type tokenRefresher struct {
	mu                 sync.Mutex
	payload            AuthenticationRequestPayload
	authenticationType AuthenticationType
	refreshAt          time.Time
}

type StoreConfig struct {
	Scheme    string
	HostName  string
	StoreCode string
}

// jwtClaims represents the minimal JWT payload we need to extract the expiry.
type jwtClaims struct {
	Exp int64 `json:"exp"`
}

// parseTokenExpiry tries to decode a JWT token and returns the time at which
// 3/4 of the token's lifetime has passed (i.e. when we should refresh).
// Returns zero time if the token cannot be parsed.
func parseTokenRefreshAt(token string) time.Time {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return time.Time{}
	}

	// JWT base64url payload — add padding if needed.
	payload := parts[1]
	switch len(payload) % 4 {
	case 2:
		payload += "=="
	case 3:
		payload += "="
	}

	decoded, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		return time.Time{}
	}

	var claims jwtClaims
	if err := json.Unmarshal(decoded, &claims); err != nil || claims.Exp == 0 {
		return time.Time{}
	}

	expiry := time.Unix(claims.Exp, 0)
	lifetime := time.Until(expiry)
	if lifetime <= 0 {
		return time.Time{}
	}

	// Refresh at 3/4 of the lifetime (i.e. subtract 1/4).
	return time.Now().Add(lifetime * 3 / 4)
}

func (c *Client) GetRouteAndDecode(route string, target interface{}, tryTo string) error {
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.WithStack(ErrNoPointer)
	}

	resp, err := c.HTTPClient.R().SetResult(target).Get(route)
	return mayReturnErrorForHTTPResponse(err, resp, tryTo)
}

func (c *Client) PostRouteAndDecode(route string, body, target interface{}, tryTo string) error {
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.WithStack(ErrNoPointer)
	}

	resp, err := c.HTTPClient.R().SetResult(target).SetBody(body).Post(route)
	return mayReturnErrorForHTTPResponse(err, resp, tryTo)
}

func NewAPIClientWithoutAuthentication(storeConfig *StoreConfig) *Client {
	httpClient := buildBasicHTTPClient(storeConfig)

	return &Client{
		HTTPClient: httpClient,
	}
}

func NewAPIClientFromAuthentication(storeConfig *StoreConfig, payload AuthenticationRequestPayload, authenticationType AuthenticationType) (*Client, error) {
	client := buildBasicHTTPClient(storeConfig)

	resp, err := client.R().SetBody(payload).Post(authenticationType.Route())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain initial auth token: %w", err)
	}

	token := mayTrimSurroundingQuotes(resp.String())
	client.SetAuthToken(token)

	// If the token is a JWT with an exp claim, set up automatic refresh.
	refreshAt := parseTokenRefreshAt(token)
	if !refreshAt.IsZero() {
		refresher := &tokenRefresher{
			payload:            payload,
			authenticationType: authenticationType,
		}
		refresher.refreshAt = refreshAt

		client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
			// Don't refresh while fetching a new token (avoid infinite loop).
			if r.URL == authenticationType.Route() {
				return nil
			}

			refresher.mu.Lock()
			defer refresher.mu.Unlock()

			// Re-check after acquiring lock — another goroutine may have already refreshed.
			if time.Now().Before(refresher.refreshAt) {
				return nil
			}

			tokenResp, err := c.R().SetBody(refresher.payload).Post(refresher.authenticationType.Route())
			if err != nil {
				return fmt.Errorf("failed to refresh auth token: %w", err)
			}

			newToken := mayTrimSurroundingQuotes(tokenResp.String())
			c.SetAuthToken(newToken)

			// Parse new token's expiry for the next refresh cycle.
			newRefreshAt := parseTokenRefreshAt(newToken)
			if !newRefreshAt.IsZero() {
				refresher.refreshAt = newRefreshAt
			} else {
				// Fallback: refresh again in 3/4 of 4 hours (Magento default).
				refresher.refreshAt = time.Now().Add(3 * time.Hour)
			}

			return nil
		})
	}

	return &Client{
		HTTPClient: client,
	}, nil
}

func NewAPIClientFromIntegration(storeConfig *StoreConfig, bearer string) (*Client, error) {
	client := buildBasicHTTPClient(storeConfig)

	client.SetAuthToken(bearer)

	return &Client{
		HTTPClient: client,
	}, nil
}

func buildBasicHTTPClient(storeConfig *StoreConfig) *resty.Client {
	apiVersion := "/V1"
	restPrefix := "/rest/" + storeConfig.StoreCode
	fullRestRoute := storeConfig.Scheme + "://" + storeConfig.HostName + restPrefix + apiVersion
	client := resty.New()
	client.SetRESTMode()
	client.SetHostURL(fullRestRoute)
	client.SetHeaders(map[string]string{
		"User-Agent": "go-magento2 (https://github.com/hermsi1337/go-magento2)",
	})

	retryWait := time.Duration(RetryWaitSeconds)
	retryMaxWait := time.Duration(RetryMaxWaitSeconds)
	client.SetRetryCount(RetryAttempts).
		SetRetryWaitTime(retryWait * time.Second).
		SetRetryMaxWaitTime(retryMaxWait * time.Second).
		AddRetryCondition(
			func(r *resty.Response) (bool, error) {
				retry := false
				status := r.StatusCode()
				if status == http.StatusServiceUnavailable || status == http.StatusInternalServerError {
					retry = true
				}
				return retry, nil
			},
		)

	return client
}
