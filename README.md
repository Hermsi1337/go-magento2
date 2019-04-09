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
    - [x] get available shipping carrier
    - [x] add shipping information (billing- and shipping-address)
    - [x] get available payment methods
    - [x] add payment method
    - [x] place order
* [x] Registered customer api
  * [x] cart
    - [x] add items
    - [x] get available shipping carrier
    - [x] add shipping information (billing- and shipping-address)
    - [x] get available payment methods
    - [x] add payment method
    - [x] place order
* [ ] Administrator

## Usage
#### Place an order as a registered customer
```go
package main

import (
	"github.com/hermsi1337/go-magento2"
	magento2Types "github.com/hermsi1337/go-magento2/types"
	"log"
)

func main() {
	// initiate storeconfig
	storeConfig := &magento2.StoreConfig{
		Scheme:    "https",
		HostName:  "localhost",
		StoreCode: "default",
	}
	// initiate login payload
	authenticationPaylod := &magento2Types.AuthenticationRequestPayload{
		Username: "anonymous@my.m2.tld",
		Password: "foobar!23",
	}
	// create a new apiClient
	apiClient, err := magento2.NewCustomerApiClient(storeConfig, authenticationPaylod)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained client: '%v'", apiClient)

	// create empty card
	cart, err := apiClient.CreateCart()
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained cart: '%v'", cart)

	// initialize items array
	var products []magento2Types.Item

	// add items to your items array
	products = append(products, magento2Types.Item{
		Sku: "123456",
		Qty: 1,
	})

	// update your cart with the desired items
	err = cart.AddItems(products)
	if err != nil {
		panic(err)
	}

	log.Printf("Added product: '%v'", cart.Detailed.Items)

	// define shipping address
	sAddr := &magento2Types.Address{
		City:      "FooCity",
		Company:   "FooCompany",
		Email:     "foo@bar.de",
		Firstname: "Foo",
		Lastname:  "Bar",
		Postcode:  "1337",
		CountryID: "DE",
		Telephone: "1337 1337 1337",
		Street:    []string{"foo", "street"},
	}

	// define billing address
	bAddr := &magento2Types.Address{
		City:      "FooCity",
		Company:   "FooCompany",
		Email:     "foo@bar.de",
		Firstname: "Foo",
		Lastname:  "Bar",
		Postcode:  "1337",
		CountryID: "DE",
		Telephone: "1337 1337 1337",
		Street:    []string{"foo", "street"},
	}

	// estimate shipping carrier for our cart
	availableCarrier, err := cart.EstimateShippingCarrier(*sAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("Obtained carrier: '%v'", availableCarrier)

	// choose your desired carrier
	desiredCarrier := availableCarrier[0]

	log.Printf("Chosen carrier: '%v'", desiredCarrier)

	// define addressinformation-payload for your cart
	payLoad := &magento2Types.AddressInformation{
		ShippingAddress:      *sAddr,
		BillingAddress:       *bAddr,
		ShippingMethodCode:   desiredCarrier.MethodCode,
		ShippingCarrierCodes: desiredCarrier.CarrierCode,
	}

	// add shipping info to cart
	err = cart.AddShippingInformation(*payLoad)
	if err != nil {
		panic(err)
	}

	// lets check what payment methods are available
	paymentMethods, err := cart.EstimatePaymentMethods()
	if err != nil {
		panic(err)
	}

	log.Printf("Obtained payment methods: '%v'", paymentMethods)

	// choose your desired payment method
	desiredPaymentMethod := paymentMethods[0]

	log.Printf("Chosen payment method: '%v'", desiredPaymentMethod)

	// create the order
	orderID, err := cart.CreateOrder(desiredPaymentMethod)
	if err != nil {
		panic(err)
	}

	// Congrats, your order has been submitted
	log.Printf("Your oder has been submitted. OrderID: '%v'", orderID)
}
```
#### Place an order as a guest
```go
package main

import (
	"github.com/hermsi1337/go-magento2"
	magento2Types "github.com/hermsi1337/go-magento2/types"
	"log"
)

func main() {
	// initiate storeconfig
	storeConfig := &magento2.StoreConfig{
		Scheme:    "https",
		HostName:  "localhost",
		StoreCode: "default",
	}
	
	// create a new apiClient
	apiClient := magento2.NewGuestApiClient(storeConfig)
	log.Printf("Obtained client: '%v'", apiClient)

	// create empty card
	cart, err := apiClient.CreateCart()
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained cart: '%v'", cart)

	// initialize items array
	var products []magento2Types.Item

	// add items to your items array
	products = append(products, magento2Types.Item{
		Sku: "161368",
		Qty: 1,
	})

	// update your cart with the desired items
	err = cart.AddItems(products)
	if err != nil {
		panic(err)
	}

	log.Printf("Added product: '%v'", cart.Detailed.Items)

	// define shipping address
	sAddr := &magento2Types.Address{
		City:      "FooCity",
		Company:   "FooCompany",
		Email:     "foo@bar.de",
		Firstname: "Foo",
		Lastname:  "Bar",
		Postcode:  "1337",
		CountryID: "DE",
		Telephone: "1337 1337 1337",
		Street:    []string{"foo", "street"},
	}

	// define billing address
	bAddr := &magento2Types.Address{
		City:      "FooCity",
		Company:   "FooCompany",
		Email:     "foo@bar.de",
		Firstname: "Foo",
		Lastname:  "Bar",
		Postcode:  "1337",
		CountryID: "DE",
		Telephone: "1337 1337 1337",
		Street:    []string{"foo", "street"},
	}

	// estimate shipping carrier for our cart
	availableCarrier, err := cart.EstimateShippingCarrier(*sAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("Obtained carrier: '%v'", availableCarrier)

	// choose your desired carrier
	desiredCarrier := availableCarrier[0]

	log.Printf("Chosen carrier: '%v'", desiredCarrier)

	// define addressinformation-payload for your cart
	payLoad := &magento2Types.AddressInformation{
		ShippingAddress:      *sAddr,
		BillingAddress:       *bAddr,
		ShippingMethodCode:   desiredCarrier.MethodCode,
		ShippingCarrierCodes: desiredCarrier.CarrierCode,
	}

	// add shipping info to cart
	err = cart.AddShippingInformation(*payLoad)
	if err != nil {
		panic(err)
	}

	// lets check what payment methods are available
	paymentMethods, err := cart.EstimatePaymentMethods()
	if err != nil {
		panic(err)
	}

	log.Printf("Obtained payment methods: '%v'", paymentMethods)

	// choose your desired payment method
	desiredPaymentMethod := paymentMethods[0]

	log.Printf("Chosen payment method: '%v'", desiredPaymentMethod)

	// create the order
	orderID, err := cart.CreateOrder(desiredPaymentMethod)
	if err != nil {
		panic(err)
	}

	// Congrats, your order has been submitted
	log.Printf("Your oder has been submitted. OrderID: '%v'", orderID)
}

```