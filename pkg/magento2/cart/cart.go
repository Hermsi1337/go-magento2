package cart

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/orders"
)

type MCart struct {
	Route     string
	QuoteID   string
	Cart      *Cart
	APIClient *api.Client
}

func NewGuestCartFromAPIClient(apiClient *api.Client) (*MCart, error) {
	mCart := &MCart{
		Cart: &Cart{},
	}
	mCart.APIClient = apiClient

	err := mCart.initializeGuestCart()
	return mCart, err
}

func NewCustomerCartFromAPIClient(apiClient *api.Client) (*MCart, error) {
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
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "initialize cart for guest")
	if err != nil {
		return err
	}

	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

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
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "initialize cart for customer")
	if err != nil {
		return err
	}

	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = customerCart
	cart.QuoteID = quoteID
	cart.APIClient = apiClient

	return cart.UpdateFromRemote()
}

func (cart *MCart) UpdateFromRemote() error {
	httpClient := cart.APIClient.HTTPClient

	resp, err := httpClient.R().SetResult(cart.Cart).Get(cart.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get detailed cart object from magento2-api")
}

func (cart *MCart) AddItems(items []Item) error {
	endpoint := cart.Route + cartItems
	httpClient := cart.APIClient.HTTPClient

	type PayLoad struct {
		CartItem Item `json:"cartItem"`
	}

	for _, item := range items {
		item.QuoteID = cart.QuoteID
		payLoad := &PayLoad{
			CartItem: item,
		}

		resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)

		err = utils.MayReturnErrorForHTTPResponse(err, resp, fmt.Sprintf("add item '%+v' to cart", item))
		if err != nil && errors.Is(err, magento2.ErrNotFound) {
			customErr := &ItemNotFoundError{ItemID: item.ItemID}
			return customErr
		} else if err != nil {
			return err
		}

		cart.Cart.Items = append(cart.Cart.Items, item)
	}

	return nil
}

func (cart *MCart) EstimateShippingCarrier(addr *Address) ([]Carrier, error) {
	endpoint := cart.Route + cartShippingCosts
	httpClient := cart.APIClient.HTTPClient

	type PayLoad struct {
		Address Address `json:"address"`
	}

	payLoad := &PayLoad{
		Address: *addr,
	}

	shippingCarrier := &[]Carrier{}

	resp, err := httpClient.R().SetBody(*payLoad).SetResult(shippingCarrier).Post(endpoint)

	return *shippingCarrier, utils.MayReturnErrorForHTTPResponse(err, resp, "estimate shipping carrier for cart")
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
	return utils.MayReturnErrorForHTTPResponse(err, resp, "add shipping information to cart")
}

func (cart *MCart) EstimatePaymentMethods() ([]PaymentMethod, error) {
	endpoint := cart.Route + cartPaymentMethods

	paymentMethods := &[]PaymentMethod{}

	err := cart.APIClient.GetRouteAndDecode(endpoint, paymentMethods, "estimate payment methods for cart")
	return *paymentMethods, err
}

func (cart *MCart) CreateOrder(paymentMethod PaymentMethod) (*orders.MOrder, error) {
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
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "create order")
	if err != nil {
		return nil, err
	}

	orderIDString := utils.MayTrimSurroundingQuotes(resp.String())
	orderIDInt, err := strconv.Atoi(orderIDString)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while extracting orderID: '%w'", err)
	}

	return &orders.MOrder{
		Route: orders.Orders + "/" + orderIDString,
		Order: &orders.Order{
			EntityID: orderIDInt,
		},
		APIClient: cart.APIClient,
	}, nil
}

func (cart *MCart) DeleteItem(itemID int) error {
	endpoint := cart.Route + cartItems + "/" + strconv.Itoa(itemID)
	httpClient := cart.APIClient.HTTPClient

	resp, err := httpClient.R().Delete(endpoint)

	return utils.MayReturnErrorForHTTPResponse(err, resp, fmt.Sprintf("delete itemID '%d'", itemID))
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
