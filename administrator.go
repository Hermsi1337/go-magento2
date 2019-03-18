package magento2

import (
	"gopkg.in/resty.v1"
)

type AdministratorApiClient struct {
	httpClient *resty.Client
}

func (client *AnonymousApiClient) AuthenticateAsAdministrator(payload AuthenticationRequestPayload) (*AdministratorApiClient, error) {
	endpoint := integrationAdminTokenService

	resp, err := client.httpClient.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.httpClient.SetAuthToken(resp.String())

	return &AdministratorApiClient{
		httpClient: client.httpClient,
	}, nil
}