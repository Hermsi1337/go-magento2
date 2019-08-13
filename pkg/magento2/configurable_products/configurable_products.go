package configurable_products

import (
	"fmt"
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
)

type MConfigurableProduct struct {
	Route     string
	Options   *[]Option
	ApiClient *api.Client
}

func SetOptionForExistingConfigurableProduct(sku string, o Option, apiClient *api.Client) (*MConfigurableProduct, error) {
	mConfigurableProduct := &MConfigurableProduct{
		Route:     configurableProducts + "/" + sku,
		Options:   &[]Option{},
		ApiClient: apiClient,
	}
	endpoint := mConfigurableProduct.Route + "/" + configurableProductsOptionsRelative
	httpClient := apiClient.HttpClient

	payLoad := createConfigurableProductByOptionPayload{
		Option: o,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)

	err = utils.MayReturnErrorForHTTPResponse(err, resp, "create configurable product")
	if err != nil {
		return mConfigurableProduct, err
	}

	err = mConfigurableProduct.UpdateOptionsFromRemote()

	return mConfigurableProduct, err
}

func (mConfigurableProduct *MConfigurableProduct) UpdateOptionsFromRemote() error {
	httpClient := mConfigurableProduct.ApiClient.HttpClient

	resp, err := httpClient.R().SetResult(mConfigurableProduct.Options).Get(mConfigurableProduct.Route + "/" + configurableProductsOptionsAllRelative)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get options for configurable product from remote")
}

func (mConfigurableProduct *MConfigurableProduct) AddChildBySKU(sku string) error {
	httpClient := mConfigurableProduct.ApiClient.HttpClient
	payLoad := addChildSKUPayload{
		Sku: sku,
	}

	endpoint := fmt.Sprintf("%s/%s", mConfigurableProduct.Route, configurableProductsChildRelative)

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "add child by sku to configurable product")
}

func GetConfigurableProductBySKU(sku string, apiClient *api.Client) (*MConfigurableProduct, error) {
	mConfigurableProduct := &MConfigurableProduct{
		Route:     configurableProducts + "/" + sku,
		Options:   &[]Option{},
		ApiClient: apiClient,
	}
	err := mConfigurableProduct.UpdateOptionsFromRemote()
	return mConfigurableProduct, err
}

func (mConfigurableProduct *MConfigurableProduct) UpdateOptionByID(o Option) error {
	httpClient := mConfigurableProduct.ApiClient.HttpClient
	endpoint := fmt.Sprintf("%s/%s/%d", mConfigurableProduct.Route, configurableProductsOptionsRelative, o.ID)

	payLoad := createConfigurableProductByOptionPayload{
		Option: o,
	}

	resp, err := httpClient.R().SetBody(payLoad).Put(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "update option for configurable product")
	if err != nil {
		return err
	}
	return mConfigurableProduct.UpdateOptionsFromRemote()
}
