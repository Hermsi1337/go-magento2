package magento2

import (
	"crypto/tls"
	"fmt"
	"github.com/hermsi1337/go-magento2/types"
	"gopkg.in/resty.v1"
)

type ApiClient struct {
	HttpClient *resty.Client
	Kind       string
}

type GuestClient interface {
	CreateCart() (Cart, error)
}

type CustomerClient interface {
	CreateCart() (Cart, error)
}

type StoreConfig struct {
	Scheme    string
	HostName  string
	StoreCode string
}

const (
	anonymousClientType     = "anonymous"
	customerClientType      = "customer"
	administratorClientType = "administrator"
)

func NewGuestApiClient(storeConfig *StoreConfig) GuestClient {
	httpClient := buildBasicHttpClient(storeConfig)

	return &ApiClient{
		HttpClient: httpClient,
		Kind:       anonymousClientType,
	}
}

func NewCustomerApiClient(storeConfig *StoreConfig, payload *types.AuthenticationRequestPayload) (CustomerClient, error) {
	client := buildBasicHttpClient(storeConfig)
	endpoint := integrationCustomerTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(mayTrimSurroundingQuotes(resp.String()))

	return &ApiClient{
		HttpClient: client,
		Kind:       customerClientType,
	}, nil
}

/*
func NewAdministratorApiClient(storeConfig *StoreConfig, payload types.AuthenticationRequestPayload) (*ApiClient, error) {
	client := buildBasicHttpClient(storeConfig)
	endpoint := integrationAdminTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &ApiClient{
		HttpClient: client,
		Kind: administratorClientType,
	}, nil
}
*/

func (apiClient *ApiClient) CreateCart() (Cart, error) {
	var cart Cart
	var endpoint string

	switch apiClient.Kind {
	case anonymousClientType:
		endpoint = guestCart
	case customerClientType:
		endpoint = customerCart
	}

	httpClient := apiClient.HttpClient
	resp, err := httpClient.R().Post(endpoint)
	if err != nil {
		return cart, err
	} else if resp.StatusCode() >= 400 {
		return cart, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}
	quoteID := mayTrimSurroundingQuotes(resp.String())

	switch apiClient.Kind {
	case anonymousClientType:
		cart.Route = guestCart + "/" + quoteID
	case customerClientType:
		cart.Route = customerCart
	}

	cart.ApiClient = apiClient
	cart.QuoteID = quoteID
	cart.Detailed, err = cart.GetDetails()
	if err != nil {
		return cart, err
	}

	return cart, err
}

func buildBasicHttpClient(storeConfig *StoreConfig) *resty.Client {
	apiVersion := "/V1"
	restPrefix := "/rest/" + storeConfig.StoreCode
	fullRestRoute := storeConfig.Scheme + "://" + storeConfig.HostName + restPrefix + apiVersion
	client := resty.New()
	client.SetRESTMode()
	client.SetHostURL(fullRestRoute)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	return client
}
