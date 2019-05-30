package place_order

import (
	"github.com/hermsi1337/go-magento2/pkg/magento2"
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
	cart, err := apiClient.NewGuestCart()
	if err != nil {
		panic(err)
	}
	log.Printf("Obtained cart: '%v'", cart)

	// initialize items array
	var products []magento2.Item

	// add items to your items array
	products = append(products, magento2.Item{
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
	sAddr := &magento2.Address{
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
	bAddr := &magento2.Address{
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
	payLoad := &magento2.AddressInformation{
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
	order, err := cart.CreateOrder(desiredPaymentMethod)
	if err != nil {
		panic(err)
	}

	// Congrats, your order has been submitted
	log.Printf("Your oder has been submitted. OrderID: '%v'", order.ID)
}
