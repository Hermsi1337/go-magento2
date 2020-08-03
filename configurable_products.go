package magento2

import (
	"fmt"
)

type MConfigurableProduct struct {
	Route     string
	Options   *[]Option
	APIClient *Client
}

func SetOptionForExistingConfigurableProduct(sku string, o *ConfigurableProductOption, apiClient *Client) (*MConfigurableProduct, error) {
	mConfigurableProduct := &MConfigurableProduct{
		Route:     configurableProducts + "/" + sku,
		Options:   &[]Option{},
		APIClient: apiClient,
	}
	endpoint := mConfigurableProduct.Route + "/" + configurableProductsOptionsRelative
	httpClient := apiClient.HTTPClient

	payLoad := createConfigurableProductByOptionPayload{
		Option: *o,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)

	err = mayReturnErrorForHTTPResponse(err, resp, "create configurable product")
	if err != nil {
		return mConfigurableProduct, err
	}

	err = mConfigurableProduct.UpdateOptionsFromRemote()

	return mConfigurableProduct, err
}

func (mConfigurableProduct *MConfigurableProduct) UpdateOptionsFromRemote() error {
	httpClient := mConfigurableProduct.APIClient.HTTPClient

	resp, err := httpClient.R().SetResult(mConfigurableProduct.Options).Get(mConfigurableProduct.Route + "/" + configurableProductsOptionsAllRelative)
	return mayReturnErrorForHTTPResponse(err, resp, "get options for configurable product from remote")
}

func (mConfigurableProduct *MConfigurableProduct) AddChildBySKU(sku string) error {
	httpClient := mConfigurableProduct.APIClient.HTTPClient
	payLoad := addChildSKUPayload{
		Sku: sku,
	}

	endpoint := fmt.Sprintf("%s/%s", mConfigurableProduct.Route, configurableProductsChildRelative)

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	return mayReturnErrorForHTTPResponse(err, resp, "add child by sku to configurable product")
}

func GetConfigurableProductBySKU(sku string, apiClient *Client) (*MConfigurableProduct, error) {
	mConfigurableProduct := &MConfigurableProduct{
		Route:     configurableProducts + "/" + sku,
		Options:   &[]Option{},
		APIClient: apiClient,
	}
	err := mConfigurableProduct.UpdateOptionsFromRemote()
	return mConfigurableProduct, err
}

func (mConfigurableProduct *MConfigurableProduct) UpdateOptionByID(o *ConfigurableProductOption) error {
	httpClient := mConfigurableProduct.APIClient.HTTPClient
	endpoint := fmt.Sprintf("%s/%s/%d", mConfigurableProduct.Route, configurableProductsOptionsRelative, o.ID)

	payLoad := createConfigurableProductByOptionPayload{
		Option: *o,
	}

	resp, err := httpClient.R().SetBody(payLoad).Put(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "update option for configurable product")
	if err != nil {
		return err
	}
	return mConfigurableProduct.UpdateOptionsFromRemote()
}
