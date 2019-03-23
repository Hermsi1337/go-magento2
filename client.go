package magento2

import (
	"fmt"
	"gopkg.in/resty.v1"
)

type ApiClient struct {
	httpClient *resty.Client
}

func NewGuestApiClient(scheme string, hostName string) *ApiClient {
	httpClient := buildBasicHttpClient(scheme, hostName)

	return &ApiClient{
		httpClient: httpClient,
	}
}

func NewCustomerApiClient(scheme string, hostName string, payload AuthenticationRequestPayload) (*ApiClient, error) {
	client := buildBasicHttpClient(scheme, hostName)
	endpoint := integrationAdminTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &ApiClient{
		httpClient: client,
	}, nil
}

func NewAdministratorApiClient(scheme string, hostName string, payload AuthenticationRequestPayload) (*ApiClient, error) {
	client := buildBasicHttpClient(scheme, hostName)
	endpoint := integrationAdminTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &ApiClient{
		httpClient: client,
	}, nil
}

func (apiClient *ApiClient) CreateCard() (Cart, error) {
	var cart Cart

	httpClient := apiClient.httpClient
	resp, err := httpClient.R().Post(guestCarts)
	if err != nil {
		return cart, err
	} else if resp.StatusCode() >= 400 {
		return cart, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	quoteID := mayTrimSurroundingQuotes(resp.String())
	cart.QuoteID = quoteID
	cart.Detailed, err = cart.GetDetails(apiClient)
	if err != nil {
		return cart, err
	}

	return cart, err
}

func buildBasicHttpClient(scheme string, hostName string) *resty.Client {
	apiVersion := "/V1"
	restPrefix := "/index.php/rest"
	fullRestRoute := scheme + "://" + hostName + restPrefix + apiVersion
	client := resty.New()
	client.SetRESTMode()
	client.SetHostURL(fullRestRoute)

	return client
}
