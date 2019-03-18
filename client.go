package magento2

import "gopkg.in/resty.v1"

func buildBasicApiClient(scheme string, hostName string) *resty.Client {
	apiVersion := "/V1"
	client := resty.New()
	client.SetRESTMode()
	client.SetHostURL(hostName + apiVersion)
	client.SetScheme(scheme)

	return client
}