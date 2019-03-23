# go-magento2
A Golang package for communicating with the magento2 api. (tested with >=2.3.0)
   
I initially built this package because I need it for a project I'm currently working on.   
Therefore I will add further features upon my need until I finished the other project.

If you need a feature which is not implemented yet, feel free to open a pull request.   
Let's make this package great together!  🚀

## Features
* [x] Guest api
  * [x] guest-carts
    - [x] add items
    - [x] estimate shipping costs
    - [x] add shipping information (billing- and shipping-address)
    - [x] add payment method
    - [x] create order
* [ ] Registered customer api
* [ ] Administrator

## Usage
Example-code for placing an order as a guest:
```go
package main

import (
	"github.com/hermsi1337/go-magento2"
	magento2Types "github.com/hermsi1337/go-magento2/types"
)

func main() {
	// create a new apiClient
	apiClient := magento2.NewGuestApiClient("http", "localhost:8080")

	// create empty card
	cart, err := apiClient.CreateGuestCard()
	if err != nil {
		panic(err)
	}

	// initialize items array
	var products []magento2Types.Item

	// add items to your items array
	products = append(products, magento2Types.Item{
		Sku: "fooduct",
		Qty: 1,
	})

	// update your cart with the desired items
	err = cart.AddItems(apiClient, products)
	if err != nil {
		panic(err)
	}

	// define shipping address
	sAddr := &magento2Types.Address{
		City: "FooCity",
		Company: "FooCompany",
		Email: "foo@bar.de",
		Firstname: "Foo",
		Lastname: "Bar",
		Postcode: "1337",
		Region: "New York",
		RegionID: 1,
		RegionCode: "NY",
		CountryID: "US",
		Telephone: "1337 1337 1337",
		Street: []string{"foo", "street"},
	}

	// define billing address
	bAddr := &magento2Types.Address{
		City: "FooCity",
		Company: "FooCompany",
		Email: "foo@bar.de",
		Firstname: "Foo",
		Lastname: "Bar",
		Postcode: "1337",
		Region: "New York",
		RegionID: 1,
		RegionCode: "NY",
		CountryID: "US",
		Telephone: "1337 1337 1337",
		Street: []string{"foo", "street"},
	}

	// estimate shipping carrier for our cart
	availableCarrier, err := cart.EstimateShippingCarrier(apiClient, *sAddr)
	if err != nil {
		panic(err)
	}

	// choose your desired carrier
	desiredCarrier := availableCarrier[0]

	// define addressinformation-payload for your cart
	payLoad := &magento2Types.AddressInformation{
		ShippingAddress: *sAddr,
		BillingAddress: *bAddr,
		ShippingMethodCode: desiredCarrier.MethodCode,
		ShippingCarrierCodes: desiredCarrier.CarrierCode,
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

	// lets check what payment methods are available
	paymentMethods, err := cart.EstimatePaymentMethods(apiClient)
	if err != nil {
		panic(err)
	}

	// choose your desired payment method
	desiredPaymentMethod := paymentMethods[0]

	// create the order
	orderID, err := cart.CreateOrder(apiClient, desiredPaymentMethod)
	if err != nil {
		panic(err)
	}

	// Congrats, your order has been submitted
	fmt.Println(orderID)
}
```