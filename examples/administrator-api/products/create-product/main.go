package main

import (
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products"
	"log"
)

func main() {
	// initiate storeconfig
	storeConfig := &api.StoreConfig{
		Scheme:    "https",
		HostName:  "magento2.hermsi.localhost",
		StoreCode: "default",
	}
	// initiate bearer payload
	bearerToken := "yd1o9zs1hb1qxnn8ek68eu8nwqjg5hrv"

	// create a new apiClient
	apiClient, err := api.NewAPIClientFromIntegration(storeConfig, bearerToken)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained client: '%v'", apiClient)

	// define your product
	product := products.Product{
		Name:           "Spaget-Shirt",
		Sku:            "spaget1234",
		Price:          1000,
		AttributeSetID: 4,
		TypeID:         "simple",
	}
	productSaveOptions := true

	// create product on remote
	mProduct, err := products.CreateOrReplaceProduct(&product, productSaveOptions, apiClient)
	if err != nil {
		panic(err)
	}

	// here you go
	log.Printf("Created product: %+v", mProduct)

	err = mProduct.UpdateQuantityForStockItem("1", 200, true)
	if err != nil {
		panic(err)
	}
}
