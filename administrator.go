package magento2

import (
	"gopkg.in/resty.v1"
)

type AdministratorApiClient struct {
	httpClient *resty.Client
}

func NewAdministratorApiClient(scheme string, hostName string, payload AuthenticationRequestPayload) (*AdministratorApiClient, error) {
	client := buildBasicApiClient(scheme, hostName)
	endpoint := integrationAdminTokenService
	resp, err := client.R().SetBody(payload).Post(endpoint)
	if err != nil {
		return nil, err
	}

	client.SetAuthToken(resp.String())

	return &AdministratorApiClient{
		httpClient: client,
	}, nil
}