package main

import (
	"log"

	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products/attribute"
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

	// define your attribute
	attr := &attribute.Attribute{
		AttributeCode:        "spagetattr",
		FrontendInput:        "select",
		DefaultFrontendLabel: "aw",
		IsRequired:           false,
	}

	// create attribute on remote
	mAttribute, err := attribute.CreateAttribute(attr, apiClient)
	if err != nil {
		panic(err)
	}

	optionValue, err := mAttribute.AddOption(attribute.Option{
		Label: "spaget",
		Value: "spaget",
	})
	if err != nil {
		panic(err)
	}

	// here you go
	log.Printf("Created attribute with ID: %+v", optionValue)
	log.Printf("Created attribute: %+v", mAttribute)
	log.Printf("Detailed attribute: %+v", mAttribute.Attribute)
}
