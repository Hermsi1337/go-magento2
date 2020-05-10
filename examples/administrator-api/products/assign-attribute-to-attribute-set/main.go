package main

import (
	"log"
	"strconv"

	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products/attribute"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products/attributeset"
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
		AttributeCode:        "awors4",
		FrontendInput:        "select",
		DefaultFrontendLabel: "aw",
		IsRequired:           false,
	}

	// create attribute on remote
	mAttribute, err := attribute.CreateAttribute(attr, apiClient)
	if err != nil {
		panic(err)
	}

	// here you go
	log.Printf("Created attribute: %+v", mAttribute)
	log.Printf("Detailed attribute: %+v", mAttribute.Attribute)

	// define your atrribute-set
	set := attributeset.AttributeSet{
		AttributeSetName: "foos4",
		SortOrder:        2,
	}

	// "Skeletonid" indicates the creation of the attribute set on the default attribute set that in Magento always has id = 4
	skeletonID := 4

	// create atrribute-set on remote
	mAttributeSet, err := attributeset.CreateAttributeSet(set, skeletonID, apiClient)
	if err != nil {
		panic(err)
	}

	// here you go
	log.Printf("Created attribute-set: %+v", mAttributeSet)
	log.Printf("Detailed attribute-set: %+v", mAttributeSet.AttributeSet)
	log.Printf("Current attribute-set attributes: %+v", mAttributeSet.AttributeSetAttributes)

	// choose in which group you want to add the attribute when assigning it to the attribute-set
	attributeGroupID, _ := strconv.Atoi(mAttributeSet.AttributeSetGroups[0].AttributeGroupID)

	// assign attribute to attribute-set
	err = mAttributeSet.AssignAttribute(attributeGroupID, 0, mAttribute.Attribute.AttributeCode)
	if err != nil {
		panic(err)
	}

	// done
	log.Printf("Updated attribute-set attributes: %+v", mAttributeSet.AttributeSetAttributes)
}
