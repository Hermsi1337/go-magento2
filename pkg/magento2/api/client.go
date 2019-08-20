package api

import (
	"crypto/tls"
	"github.com/hermsi1337/go-magento2/internal/utils"
	"gopkg.in/resty.v1"
	"time"
)

type Client struct {
	HttpClient *resty.Client
	Kind       string
}

type StoreConfig struct {
	Scheme    string
	HostName  string
	StoreCode string
}

func NewGuestApiClient(storeConfig *StoreConfig) *Client {
	httpClient := buildBasicHttpClient(storeConfig)

	return &Client{
		HttpClient: httpClient,
		Kind:       AnonymousClientType,
	}
}

func NewCustomerApiClient(storeConfig *StoreConfig, payload AuthenticationRequestPayload) (*Client, error) {
	client := buildBasicHttpClient(storeConfig)
	endpoint := integrationCustomerTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(utils.MayTrimSurroundingQuotes(resp.String()))

	return &Client{
		HttpClient: client,
		Kind:       CustomerClientType,
	}, nil
}

func NewAdministratorApiClientFromAuthentication(storeConfig *StoreConfig, payload AuthenticationRequestPayload) (*Client, error) {
	client := buildBasicHttpClient(storeConfig)
	endpoint := integrationAdminTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &Client{
		HttpClient: client,
		Kind:       AdministratorClientType,
	}, nil
}

func NewAdministratorApiClientFromIntegration(storeConfig *StoreConfig, bearer string) (*Client, error) {
	client := buildBasicHttpClient(storeConfig)

	client.SetAuthToken(bearer)

	return &Client{
		HttpClient: client,
		Kind:       AdministratorClientType,
	}, nil
}

func buildBasicHttpClient(storeConfig *StoreConfig) *resty.Client {
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
	client.SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second)

	return client
}
