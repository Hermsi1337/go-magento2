package products

import (
	"fmt"
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
)

const (
	products = "/products"
)

type MProduct struct {
	Route     string
	Product   Product
	ApiClient *api.Client
}

func CreateOrReplaceProduct(product Product, saveOptions bool, apiClient *api.Client) (*MProduct, error) {
	mp := &MProduct{
		Product:   product,
		ApiClient: apiClient,
	}

	err := mp.createOrReplaceProduct(saveOptions)

	return mp, err
}

func GetProductBySKU(sku string, apiClient *api.Client) (*MProduct, error) {
	mp := &MProduct{
		Route:     products + "/" + sku,
		ApiClient: apiClient,
	}

	err := mp.updateProductFromRemote()

	return mp, err
}

func (mProduct *MProduct) createOrReplaceProduct(saveOptions bool) error {
	mp := &MProduct{}
	endpoint := products
	httpClient := mProduct.ApiClient.HttpClient

	payLoad := AddProductPayload{
		Product:     mProduct.Product,
		SaveOptions: saveOptions,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(&mp.Product).Post(endpoint)
	if err != nil {
		return fmt.Errorf("received error while creating new product '%v': '%v' | payload: '%+v' ", mProduct.Product, err, payLoad)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v' - body: '%+v'", resp.StatusCode(), resp, resp.Request.Body)
	}

	productSKU := utils.MayTrimSurroundingQuotes(mp.Product.Sku)
	mProduct.Route = products + "/" + productSKU

	return nil
}

func (mProduct *MProduct) updateProductFromRemote() error {
	httpClient := mProduct.ApiClient.HttpClient

	resp, err := httpClient.R().SetResult(mProduct.Product).Get(mProduct.Route)
	if err != nil {
		return fmt.Errorf("error while getting detailed product from magento2-api: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	return nil
}
