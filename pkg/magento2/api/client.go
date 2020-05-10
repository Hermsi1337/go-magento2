package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/pkg/errors"

	"github.com/hermsi1337/go-magento2/internal/utils"
	"gopkg.in/resty.v1"
)

const (
	RetryAttemps        = 3
	RetryWaitSeconds    = 5
	RetryMaxWaitSeconds = 20
)

type Client struct {
	HTTPClient *resty.Client
}

type StoreConfig struct {
	Scheme    string
	HostName  string
	StoreCode string
}

var ErrNoPointer = fmt.Errorf("target interface must be a pointer")

func (c *Client) GetRouteAndDecode(route string, target interface{}, tryTo string) error {
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.WithStack(ErrNoPointer)
	}

	resp, err := c.HTTPClient.R().SetResult(target).Get(route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, tryTo)
}

func (c *Client) PostRouteAndDecode(route string, body, target interface{}, tryTo string) error {
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.WithStack(ErrNoPointer)
	}

	resp, err := c.HTTPClient.R().SetResult(target).SetBody(body).Post(route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, tryTo)
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
		return nil, err
	}

	client.SetAuthToken(utils.MayTrimSurroundingQuotes(resp.String()))

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
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: false})
	client.SetHeaders(map[string]string{
		"User-Agent": "go-magento2 (https://github.com/hermsi1337/go-magento2)",
	})

	retryWait := time.Duration(RetryWaitSeconds)
	retryMaxWait := time.Duration(RetryMaxWaitSeconds)
	client.SetRetryCount(RetryAttemps).
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
