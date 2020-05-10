package products

import (
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
)

const (
	products = "/products"
)

type MProduct struct {
	Route     string
	Product   *Product
	APIClient *api.Client
}

func CreateOrReplaceProduct(product *Product, saveOptions bool, apiClient *api.Client) (*MProduct, error) {
	mp := &MProduct{
		Product:   product,
		APIClient: apiClient,
	}

	err := mp.createOrReplaceProduct(saveOptions)

	return mp, err
}

func GetProductBySKU(sku string, apiClient *api.Client) (*MProduct, error) {
	mProduct := &MProduct{
		Route:     products + "/" + sku,
		Product:   &Product{},
		APIClient: apiClient,
	}

	err := mProduct.UpdateProductFromRemote()

	return mProduct, err
}

func (mProduct *MProduct) createOrReplaceProduct(saveOptions bool) error {
	endpoint := products
	httpClient := mProduct.APIClient.HTTPClient

	payLoad := AddProductPayload{
		Product:     *mProduct.Product,
		SaveOptions: saveOptions,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mProduct.Product).Post(endpoint)
	productSKU := utils.MayTrimSurroundingQuotes(mProduct.Product.Sku)
	mProduct.Route = products + "/" + productSKU

	return utils.MayReturnErrorForHTTPResponse(err, resp, "create new product on remote")
}

func (mProduct *MProduct) UpdateProductFromRemote() error {
	httpClient := mProduct.APIClient.HTTPClient

	resp, err := httpClient.R().SetResult(mProduct.Product).Get(mProduct.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get detailed product from remote")
}

func (mProduct *MProduct) UpdateQuantityForStockItem(stockItem string, quantity int, isInStock bool) error {
	httpClient := mProduct.APIClient.HTTPClient

	updateStockPayload := updateStockPayload{StockItem: StockItem{Qty: quantity, IsInStock: isInStock}}

	resp, err := httpClient.R().SetBody(updateStockPayload).Put(mProduct.Route + "/" + stockItemsRelative + "/" + stockItem)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "update stock for product")
}
