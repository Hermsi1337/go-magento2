package magento2

import "fmt"

type MProduct struct {
	Route string
	Product Product
	ApiClient *ApiClient
}

func (apiClient *ApiClient) CreateOrUpdateProduct(p Product, saveOptions bool) (*MProduct, error) {
	mp := &MProduct{}
	endpoint := products
	httpClient := apiClient.HttpClient

	payLoad := AddProductPayload{
		Product: p,
		SaveOptions: saveOptions,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(&mp.Product).Post(endpoint)
	if err != nil {
		return nil, fmt.Errorf("received error while creating new product '%v': '%v' | payload: '%+v' ", p, err, payLoad)
	} else if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("unexpected statuscode '%v' - response: '%v' - body: '%+v'", resp.StatusCode(), resp, resp.Request.Body)
	}

	return mp, nil
}