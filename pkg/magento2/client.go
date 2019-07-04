package magento2

import (
	"crypto/tls"
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
	CreateOrUpdateProduct(p Product, saveOptions bool) (*MProduct, error)
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
