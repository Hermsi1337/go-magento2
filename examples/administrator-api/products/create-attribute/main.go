package main

import (
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products/attribute"
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
	apiClient, err := api.NewAdministratorApiClientFromIntegration(storeConfig, bearerToken)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained client: '%v'", apiClient)

	// define your atrribute
	attr := attribute.Attribute{
		AttributeCode: "awor",
		FrontendInput: "select",
		DefaultFrontendLabel: "aw",
		IsRequired: false,
	}
	// create atrribute-set on remote
	mAttribute, err := attribute.CreateAttribute(attr, apiClient)
	if err != nil {
		panic(err)
	}

	// here you go
	log.Printf("Created attribute: %+v", mAttribute)
	log.Printf("Detailed attribute: %+v", mAttribute.Attribute)
}
