package magento2

import (
	"gopkg.in/resty.v1"
)

type GuestApiClient struct {
	httpClient *resty.Client
}

func NewGuestApiClient(scheme string, hostName string) *GuestApiClient {
	client := buildBasicApiClient(scheme, hostName)

	return &GuestApiClient{
		httpClient: client,
	}
}