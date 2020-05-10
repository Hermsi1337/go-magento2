package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/configurableproducts"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products/attribute"
)

// TODO: FINISH EXAMPLE

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
		AttributeCode:        "myselectattribute",
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

	optionValueInt, err := strconv.Atoi(optionValue)
	if err != nil {
		panic(err)
	}

	option := &configurableproducts.Option{
		AttributeID:  fmt.Sprintf("%d", mAttribute.Attribute.AttributeID),
		Label:        mAttribute.Attribute.DefaultFrontendLabel,
		Position:     0,
		IsUseDefault: false,
		Values: []configurableproducts.Value{
			{
				ValueIndex: optionValueInt,
			},
		},
	}

	mOption, err := configurableproducts.SetOptionForExistingConfigurableProduct("configurableSpaget", option, apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created configurable Product: '%+v'", mOption)
	fmt.Printf("Created configurable Product options: '%+v'", mOption.Options)
}
