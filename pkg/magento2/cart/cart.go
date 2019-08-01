package cart

import (
	"fmt"
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/orders"
	"strconv"
)

type MCart struct {
	Route     string
	QuoteID   string
	Cart      *Cart
	ApiClient *api.Client
}

func NewGuestCartFromApiClient(apiClient *api.Client) (*MCart, error) {
	mCart := &MCart{
		Cart: &Cart{},
	}
	mCart.ApiClient = apiClient

	err := mCart.initializeGuestCart()
	return mCart, err
}

func NewCustomerCartFromApiClient(apiClient *api.Client) (*MCart, error) {
	mCart := &MCart{
		Cart: &Cart{},
	}
	mCart.ApiClient = apiClient

	err := mCart.initializeCustomerCart()
	return mCart, err
}

func (cart *MCart) initializeGuestCart() error {
	endpoint := guestCart
	apiClient := cart.ApiClient

	httpClient := apiClient.HttpClient
	resp, err := httpClient.R().Post(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "initialize cart for guest")
	if err != nil {
		return err
	}

	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = guestCart + "/" + quoteID
	cart.QuoteID = quoteID
	cart.ApiClient = apiClient

	return cart.UpdateCartFromRemote()
}

func (cart *MCart) initializeCustomerCart() error {
	endpoint := customerCart
	apiClient := cart.ApiClient

	httpClient := apiClient.HttpClient
	resp, err := httpClient.R().Post(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "initialize cart for customer")
	if err != nil {
		return err
	}

	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = customerCart
	cart.QuoteID = quoteID
	cart.ApiClient = apiClient

	return cart.UpdateCartFromRemote()
}

func (cart *MCart) UpdateCartFromRemote() error {
	httpClient := cart.ApiClient.HttpClient

	resp, err := httpClient.R().SetResult(cart.Cart).Get(cart.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get detailed cart object from magento2-api")
}

func (cart *MCart) AddItems(items []Item) error {
	endpoint := cart.Route + cartItems
	httpClient := cart.ApiClient.HttpClient

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
		if err != nil {
			return err
		}
		cart.Cart.Items = append(cart.Cart.Items, item)
	}

	return nil
}

func (cart *MCart) EstimateShippingCarrier(addrInfo Address) ([]Carrier, error) {
	endpoint := cart.Route + cartShippingCosts
	httpClient := cart.ApiClient.HttpClient

	type PayLoad struct {
		Address Address `json:"address"`
	}

	payLoad := &PayLoad{
		Address: addrInfo,
	}

	shippingCarrier := &[]Carrier{}

	resp, err := httpClient.R().SetBody(*payLoad).SetResult(shippingCarrier).Post(endpoint)

	return *shippingCarrier, utils.MayReturnErrorForHTTPResponse(err, resp, "estimate shipping carrier for cart")
}

func (cart *MCart) AddShippingInformation(addrInfo AddressInformation) error {
	endpoint := cart.Route + cartShippingInformation
	httpClient := cart.ApiClient.HttpClient

	type PayLoad struct {
		AddressInformation AddressInformation `json:"addressInformation"`
	}

	payLoad := &PayLoad{
		AddressInformation: addrInfo,
	}

	resp, err := httpClient.R().SetBody(*payLoad).Post(endpoint)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "add shipping information to cart")
}

func (cart *MCart) EstimatePaymentMethods() ([]PaymentMethod, error) {
	endpoint := cart.Route + cartPaymentMethods
	httpClient := cart.ApiClient.HttpClient

	paymentMethods := &[]PaymentMethod{}

	resp, err := httpClient.R().SetResult(paymentMethods).Get(endpoint)

	return *paymentMethods, utils.MayReturnErrorForHTTPResponse(err, resp, "estimate payment methods for cart")
}

func (cart *MCart) CreateOrder(paymentMethod PaymentMethod) (*orders.MOrder, error) {
	endpoint := cart.Route + cartPlaceOrder
	httpClient := cart.ApiClient.HttpClient

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
		return nil, fmt.Errorf("unexpected error while extracting orderID: '%v'", err)
	}

	return &orders.MOrder{
		Route:   orders.Orders + "/" + orderIDString,
		Order: &orders.Order{
			EntityID: orderIDInt,
		},
		ApiClient: cart.ApiClient,
	}, nil
}

func (cart *MCart) DeleteItem(itemID int) error {
	endpoint := cart.Route + cartItems + "/" + strconv.Itoa(itemID)
	httpClient := cart.ApiClient.HttpClient

	resp, err := httpClient.R().Delete(endpoint)

	return utils.MayReturnErrorForHTTPResponse(err, resp, fmt.Sprintf("delete itemID '%d'", itemID))
}

func (cart *MCart) DeleteAllItems() error {
	err := cart.UpdateCartFromRemote()
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
