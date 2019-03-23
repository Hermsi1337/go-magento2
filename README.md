# go-magento2
A Golang package for communicating with the magento2 api. (tested with >=2.3.0)
   
I initially built this package because I need it for a project I'm currently working on.   
Therefore it is in a very early state and under steady development. Pull requests are very welcome!

## Next up
* Guest api
  * checkout
    - [ ] Add desired payment method
    - [ ] Place order
* Registered customer api
* Administrator

## Usage
Up until today (23. March 2019) the package only supports several features of the [Magento2-guest-api](https://devdocs.magento.com/redoc/2.3/guest-rest-api.html).   

Example-code for creating a guest-cart and adding items as well as shipping-information to it:
```go
package main

import (
	"github.com/hermsi1337/go-magento2"
)

func main() {
	// create a new apiClient
	apiClient := magento2.NewGuestApiClient("http", "localhost:8080")

	// create empty card
	cart, err := apiClient.CreateCard()
	if err != nil {
		panic(err)
	}

	// initialize items array
	var products []magento2.Item
	
	// add items to your items array
	products = append(products, magento2.Item{
		Sku: "fooduct",
		Qty: 1,
	})

	// update your cart with the desired items
	err = cart.AddItems(apiClient, products)
	if err != nil {
		panic(err)
	}

	// define shipping address
	sAddr := &magento2.Address{
		City: "noida",
		Company: "ipgrad",
		Email: "fasf@fasdf.de",
		Firstname: "troll",
		Lastname: "trroll",
		Postcode: "32345",
		Region: "New York",
		RegionID: 0,
		RegionCode: "NY",
		CountryID: "US",
		Telephone: "3414124312413",
		Street: []string{"fsafas", "fasfasfd"},
	}
	
	// define billing address
	bAddr := &magento2.Address{
		City: "noida",
		Company: "ipgrad",
		Email: "fasf@fasdf.de",
		Firstname: "troll",
		Lastname: "trroll",
		Postcode: "32345",
		Region: "New York",
		RegionID: 0,
		RegionCode: "NY",
		CountryID: "US",
		Telephone: "3414124312413",
		Street: []string{"fsafas", "fasfasfd"},
	}
	
	// define addressinformation-payload for your cart
	payLoad := &magento2.AddressInformation{
		ShippingAddress: *sAddr,
		BillingAddress: *bAddr,
		ShippingMethodCode: "flatrate",
		ShippingCarrierCodes: "flatrate",
	}

	// add shipping info to cart
	err = cart.AddShippingInformation(apiClient, *payLoad)
	if err != nil {
		panic(err)
	}
	
	// update your cart object in order to get the current state from magento2 api
	err = cart.UpdateSelf(apiClient)
	if err != nil {
		panic(err)
	}
}
```