package main

import (
	"fmt"
	"log"

	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/categories"
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

	c := &categories.Category{
		Name:     "spagetegory",
		Level:    2,
		IsActive: true,
	}

	mC, err := categories.CreateCategory(c, apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n MagentoCategory struct: '%+v'", mC)
	fmt.Printf("\n MagentoCategory remote: %+v", mC.Category)
}
