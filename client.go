package magento2

import "gopkg.in/resty.v1"

func buildBasicApiClient(scheme string, hostName string) *resty.Client {
	apiVersion := "/V1"
	restPrefix := "/index.php/rest"
	fullRestRoute := hostName + restPrefix + apiVersion
	client := resty.New()
	client.SetRESTMode()
	client.SetHostURL(fullRestRoute)
	client.SetScheme(scheme)

	return client
}