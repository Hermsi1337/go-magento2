package magento2

const (
	products = "/products"
)

type MProduct struct {
	Route     string
	Product   *Product
	APIClient *Client
}

func CreateOrReplaceProduct(product *Product, saveOptions bool, apiClient *Client) (*MProduct, error) {
	mp := &MProduct{
		Product:   product,
		APIClient: apiClient,
	}

	err := mp.createOrReplaceProduct(saveOptions)

	return mp, err
}

func GetProductBySKU(sku string, apiClient *Client) (*MProduct, error) {
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
	productSKU := mayTrimSurroundingQuotes(mProduct.Product.Sku)
	mProduct.Route = products + "/" + productSKU

	return mayReturnErrorForHTTPResponse(err, resp, "create new product on remote")
}

func (mProduct *MProduct) UpdateProductFromRemote() error {
	httpClient := mProduct.APIClient.HTTPClient

	resp, err := httpClient.R().SetResult(mProduct.Product).Get(mProduct.Route)
	return mayReturnErrorForHTTPResponse(err, resp, "get detailed product from remote")
}

func (mProduct *MProduct) UpdateQuantityForStockItem(stockItem string, quantity int, isInStock bool) error {
	httpClient := mProduct.APIClient.HTTPClient

	updateStockPayload := updateStockPayload{StockItem: StockItem{Qty: quantity, IsInStock: isInStock}}

	resp, err := httpClient.R().SetBody(updateStockPayload).Put(mProduct.Route + "/" + stockItemsRelative + "/" + stockItem)
	return mayReturnErrorForHTTPResponse(err, resp, "update stock for product")
}
