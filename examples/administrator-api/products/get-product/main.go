package main

import (
	"github.com/hermsi1337/go-magento2"
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
	apiClient, err := magento2.NewAPIClientFromIntegration(storeConfig, bearerToken)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained client: '%v'", apiClient)

	// get product by SKU
	mProduct, err := magento2.GetProductBySKU("spaget1234", apiClient)
	if err != nil {
		panic(err)
	}

	// here you go
	log.Printf("Got product: %+v", mProduct)
}
