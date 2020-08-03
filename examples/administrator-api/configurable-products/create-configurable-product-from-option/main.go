package main

import (
	"fmt"
	"github.com/hermsi1337/go-magento2"
	"log"
	"strconv"
)

// TODO: FINISH EXAMPLE

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

	// define your attribute
	attr := &magento2.Attribute{
		AttributeCode:        "myselectattribute",
		FrontendInput:        "select",
		DefaultFrontendLabel: "aw",
		IsRequired:           false,
	}

	// create attribute on remote
	mAttribute, err := magento2.CreateAttribute(attr, apiClient)
	if err != nil {
		panic(err)
	}

	optionValue, err := mAttribute.AddOption(magento2.Option{
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

	option := &magento2.ConfigurableProductOption{
		AttributeID:  fmt.Sprintf("%d", mAttribute.Attribute.AttributeID),
		Label:        mAttribute.Attribute.DefaultFrontendLabel,
		Position:     0,
		IsUseDefault: false,
		Values: []magento2.Value{
			{
				ValueIndex: optionValueInt,
			},
		},
	}

	mOption, err := magento2.SetOptionForExistingConfigurableProduct("configurableSpaget", option, apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created configurable Product: '%+v'", mOption)
	fmt.Printf("Created configurable Product options: '%+v'", mOption.Options)
}
