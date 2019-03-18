package magento2

import (
	"gopkg.in/resty.v1"
)

type AnonymousApiClient struct {
	httpClient *resty.Client
}

func NewAnonymousApiClient(scheme string, hostName string) *AnonymousApiClient {
	apiVersion := "/V1"
	client := resty.New()
	client.SetRESTMode()
	client.SetHostURL(hostName + apiVersion)
	client.SetScheme(scheme)

	return &AnonymousApiClient{
		httpClient: client,
	}
}