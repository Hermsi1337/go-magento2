package main

import (
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/cart"
	"log"
)

func main() {
	// initiate storeconfig
	storeConfig := &api.StoreConfig{
		Scheme:    "https",
		HostName:  "magento2.hermsi.localhost",
		StoreCode: "default",
	}

	// initiate login payload
	authenticationPaylod := api.AuthenticationRequestPayload{
		Username: "anonymous@my.m2.tld",
		Password: "foobar!23",
	}

	// create a new apiClient
	authenticationType := api.Customer
	apiClient, err := api.NewAPIClientFromAuthentication(storeConfig, authenticationPaylod, authenticationType)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained client: '%+v'", apiClient)

	// create empty card
	mCart, err := cart.NewCustomerCartFromAPIClient(apiClient)
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained mCart: '%+v'", mCart)
	log.Printf("Detailed cart: '%+v'", mCart.Cart)

	// initialize items array
	var products []cart.Item

	// add items to your items array
	products = append(products, cart.Item{
		Sku: "123456",
		Qty: 1,
	})

	// update your cart with the desired items
	err = mCart.AddItems(products)
	if err != nil {
		panic(err)
	}

	log.Printf("Added products: '%+v'", products)
	log.Printf("Products in cart: '%+v'", mCart.Cart.Items)

	// define shipping address
	sAddr := &cart.Address{
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
	bAddr := &cart.Address{
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
	availableCarrier, err := mCart.EstimateShippingCarrier(sAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("Obtained carrier: '%+v'", availableCarrier)

	// choose your desired carrier
	desiredCarrier := availableCarrier[0]

	log.Printf("Chosen carrier: '%+v'", desiredCarrier)

	// define addressinformation-payload for your cart
	payLoad := &cart.AddressInformation{
		ShippingAddress:      *sAddr,
		BillingAddress:       *bAddr,
		ShippingMethodCode:   desiredCarrier.MethodCode,
		ShippingCarrierCodes: desiredCarrier.CarrierCode,
	}

	// add shipping info to cart
	err = mCart.AddShippingInformation(payLoad)
	if err != nil {
		panic(err)
	}

	// lets check what payment methods are available
	paymentMethods, err := mCart.EstimatePaymentMethods()
	if err != nil {
		panic(err)
	}

	log.Printf("Obtained payment methods: '%+v'", paymentMethods)

	// choose your desired payment method
	desiredPaymentMethod := paymentMethods[0]

	log.Printf("Chosen payment method: '%+v'", desiredPaymentMethod)

	// create the order
	order, err := mCart.CreateOrder(desiredPaymentMethod)
	if err != nil {
		panic(err)
	}

	// Congrats, your order has been submitted
	log.Printf("Your oder has been submitted. OrderID: '%v'", order.Order.EntityID)
}
