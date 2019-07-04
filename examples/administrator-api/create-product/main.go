package main

import (
	"github.com/hermsi1337/go-magento2/pkg/magento2"
	"log"
)

func main() {
	// initiate storeconfig
	storeConfig := &magento2.StoreConfig{
		Scheme:    "https",
		HostName:  "magento2.hermsi.localhost",
		StoreCode: "default",
	}
	// initiate bearer payload
	bearerToken := "yd1o9zs1hb1qxnn8ek68eu8nwqjg5hrv"

	// create a new apiClient
	apiClient, err := magento2.NewAdministratorApiClientFromIntegration(storeConfig, bearerToken)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained client: '%v'", apiClient)

	product := magento2.Product{
		Name: "Spaget-Shirt",
		Sku: "spaget1234",
		Price: 100,
		AttributeSetID: 4,
		TypeID: "simple",
	}

	mProduct, err := apiClient.CreateOrUpdateProduct(product, false)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", mProduct)
}