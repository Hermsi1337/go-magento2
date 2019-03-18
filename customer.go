package magento2

import "gopkg.in/resty.v1"

type CustomerApiClient struct {
	httpClient *resty.Client
}

func (client *AnonymousApiClient) AuthenticateAsCustomer(payload AuthenticationRequestPayload) (*CustomerApiClient, error) {
	endpoint := integrationCustomerTokenService

	resp, err := client.httpClient.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.httpClient.SetAuthToken(resp.String())

	return &CustomerApiClient{
		httpClient: client.httpClient,
	}, nil
}