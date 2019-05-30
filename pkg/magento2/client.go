package magento2

import (
	"crypto/tls"
	"fmt"
	"github.com/hermsi1337/go-magento2/internal/utils"

	"gopkg.in/resty.v1"
)

type ApiClient struct {
	HttpClient *resty.Client
	Kind       string
}

type GuestClient interface {
	NewGuestCart() (Cart, error)
}

type CustomerClient interface {
	NewCustomerCart() (Cart, error)
}

type AdministratorClient interface {
	NewGuestCart() (Cart, error)
	NewCustomerCart() (Cart, error)
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

func NewCustomerApiClient(storeConfig *StoreConfig, payload *AuthenticationRequestPayload) (CustomerClient, error) {
	client := buildBasicHttpClient(storeConfig)
	endpoint := integrationCustomerTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(utils.MayTrimSurroundingQuotes(resp.String()))

	return &ApiClient{
		HttpClient: client,
		Kind:       customerClientType,
	}, nil
}

func NewAdministratorApiClientFromAuthentication(storeConfig *StoreConfig, payload AuthenticationRequestPayload) (AdministratorClient, error) {
	client := buildBasicHttpClient(storeConfig)
	endpoint := integrationAdminTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &ApiClient{
		HttpClient: client,
		Kind:       administratorClientType,
	}, nil
}

func NewAdministratorApiClientFromIntegration(storeConfig *StoreConfig, bearer string) (AdministratorClient, error) {
	client := buildBasicHttpClient(storeConfig)

	client.SetAuthToken(bearer)

	return &ApiClient{
		HttpClient: client,
		Kind:       administratorClientType,
	}, nil
}

func (apiClient *ApiClient) NewGuestCart() (Cart, error) {
	var cart Cart
	endpoint := guestCart

	httpClient := apiClient.HttpClient
	resp, err := httpClient.R().Post(endpoint)
	if err != nil {
		return cart, err
	} else if resp.StatusCode() >= 400 {
		return cart, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}
	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = guestCart + "/" + quoteID

	cart.ApiClient = apiClient
	cart.QuoteID = quoteID
	cart.Detailed, err = cart.GetDetails()
	if err != nil {
		return cart, err
	}

	return cart, err
}

func (apiClient *ApiClient) NewCustomerCart() (Cart, error) {
	var cart Cart
	endpoint := customerCart

	httpClient := apiClient.HttpClient
	resp, err := httpClient.R().Post(endpoint)
	if err != nil {
		return cart, err
	} else if resp.StatusCode() >= 400 {
		return cart, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}
	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = customerCart

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
