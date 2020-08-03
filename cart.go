package magento2

import (
	"errors"
	"fmt"
	"strconv"
)

type MCart struct {
	Route     string
	QuoteID   string
	Cart      *Cart
	APIClient *Client
}

func NewGuestCartFromAPIClient(apiClient *Client) (*MCart, error) {
	mCart := &MCart{
		Cart: &Cart{},
	}
	mCart.APIClient = apiClient

	err := mCart.initializeGuestCart()
	return mCart, err
}

func NewCustomerCartFromAPIClient(apiClient *Client) (*MCart, error) {
	mCart := &MCart{
		Cart: &Cart{},
	}
	mCart.APIClient = apiClient

	err := mCart.initializeCustomerCart()
	return mCart, err
}

func (cart *MCart) initializeGuestCart() error {
	endpoint := guestCart
	apiClient := cart.APIClient

	httpClient := apiClient.HTTPClient
	resp, err := httpClient.R().Post(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "initialize cart for guest")
	if err != nil {
		return err
	}

	quoteID := mayTrimSurroundingQuotes(resp.String())

	cart.Route = guestCart + "/" + quoteID
	cart.QuoteID = quoteID
	cart.APIClient = apiClient

	return cart.UpdateFromRemote()
}

func (cart *MCart) initializeCustomerCart() error {
	endpoint := customerCart
	apiClient := cart.APIClient

	httpClient := apiClient.HTTPClient
	resp, err := httpClient.R().Post(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "initialize cart for customer")
	if err != nil {
		return err
	}

	quoteID := mayTrimSurroundingQuotes(resp.String())

	cart.Route = customerCart
	cart.QuoteID = quoteID
	cart.APIClient = apiClient

	return cart.UpdateFromRemote()
}

func (cart *MCart) UpdateFromRemote() error {
	httpClient := cart.APIClient.HTTPClient

	resp, err := httpClient.R().SetResult(cart.Cart).Get(cart.Route)
	return mayReturnErrorForHTTPResponse(err, resp, "get detailed cart object from magento2-api")
}

func (cart *MCart) AddItems(items []CartItem) error {
	endpoint := cart.Route + cartItems
	httpClient := cart.APIClient.HTTPClient

	type PayLoad struct {
		CartItem CartItem `json:"cartItem"`
	}

	for _, item := range items {
		item.QuoteID = cart.QuoteID
		payLoad := &PayLoad{
			CartItem: item,
		}

		resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)

		err = mayReturnErrorForHTTPResponse(err, resp, fmt.Sprintf("add item '%+v' to cart", item))
		if err != nil && errors.Is(err, ErrNotFound) {
			customErr := &ItemNotFoundError{ItemID: item.ItemID}
			return customErr
		} else if err != nil {
			return err
		}

		cart.Cart.Items = append(cart.Cart.Items, item)
	}

	return nil
}

func (cart *MCart) EstimateShippingCarrier(addr *ShippingAddress) ([]Carrier, error) {
	endpoint := cart.Route + cartShippingCosts
	httpClient := cart.APIClient.HTTPClient

	type PayLoad struct {
		Address ShippingAddress `json:"address"`
	}

	payLoad := &PayLoad{
		Address: *addr,
	}

	shippingCarrier := &[]Carrier{}

	resp, err := httpClient.R().SetBody(*payLoad).SetResult(shippingCarrier).Post(endpoint)

	return *shippingCarrier, mayReturnErrorForHTTPResponse(err, resp, "estimate shipping carrier for cart")
}

func (cart *MCart) AddShippingInformation(addrInfo *AddressInformation) error {
	endpoint := cart.Route + cartShippingInformation
	httpClient := cart.APIClient.HTTPClient

	type PayLoad struct {
		AddressInformation AddressInformation `json:"addressInformation"`
	}

	payLoad := &PayLoad{
		AddressInformation: *addrInfo,
	}

	resp, err := httpClient.R().SetBody(*payLoad).Post(endpoint)
	return mayReturnErrorForHTTPResponse(err, resp, "add shipping information to cart")
}

func (cart *MCart) EstimatePaymentMethods() ([]PaymentMethod, error) {
	endpoint := cart.Route + cartPaymentMethods

	paymentMethods := &[]PaymentMethod{}

	err := cart.APIClient.GetRouteAndDecode(endpoint, paymentMethods, "estimate payment methods for cart")
	return *paymentMethods, err
}

func (cart *MCart) CreateOrder(paymentMethod PaymentMethod) (*MOrder, error) {
	endpoint := cart.Route + cartPlaceOrder
	httpClient := cart.APIClient.HTTPClient

	type PayLoad struct {
		PaymentMethod PaymentMethodCode `json:"paymentMethod"`
	}

	payLoad := &PayLoad{
		PaymentMethod: PaymentMethodCode{
			Method: paymentMethod.Code,
		},
	}

	resp, err := httpClient.R().SetBody(payLoad).Put(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "create order")
	if err != nil {
		return nil, err
	}

	orderIDString := mayTrimSurroundingQuotes(resp.String())
	orderIDInt, err := strconv.Atoi(orderIDString)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while extracting orderID: '%w'", err)
	}

	return &MOrder{
		Route: Orders + "/" + orderIDString,
		Order: &Order{
			EntityID: orderIDInt,
		},
		APIClient: cart.APIClient,
	}, nil
}

func (cart *MCart) DeleteItem(itemID int) error {
	endpoint := cart.Route + cartItems + "/" + strconv.Itoa(itemID)
	httpClient := cart.APIClient.HTTPClient

	resp, err := httpClient.R().Delete(endpoint)

	return mayReturnErrorForHTTPResponse(err, resp, fmt.Sprintf("delete itemID '%d'", itemID))
}

func (cart *MCart) DeleteAllItems() error {
	err := cart.UpdateFromRemote()
	if err != nil {
		return err
	}

	for i := range cart.Cart.Items {
		err = cart.DeleteItem(cart.Cart.Items[i].ItemID)
		if err != nil {
			return err
		}
	}

	return nil
}
