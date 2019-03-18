package magento2

import (
	"gopkg.in/resty.v1"
)

type CustomerApiClient struct {
	httpClient *resty.Client
}

func NewCustomerApiClient(scheme string, hostName string, payload AuthenticationRequestPayload) (*CustomerApiClient, error) {
	client := buildBasicApiClient(scheme, hostName)
	endpoint := integrationCustomerTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &CustomerApiClient{
		httpClient: client,
	}, nil
}