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
	if err != nil {
		return err
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}
	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = guestCart + "/" + quoteID

	cart.ApiClient = apiClient
	cart.QuoteID = quoteID

	return cart.updateCartFromRemote()
}

func (cart *MCart) initializeCustomerCart() error{
	endpoint := customerCart
	apiClient := cart.ApiClient

	httpClient := apiClient.HttpClient
	resp, err := httpClient.R().Post(endpoint)
	if err != nil {
		return err
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}
	quoteID := utils.MayTrimSurroundingQuotes(resp.String())

	cart.Route = customerCart

	cart.ApiClient = apiClient
	cart.QuoteID = quoteID

	return cart.updateCartFromRemote()
}

func (cart *MCart) updateCartFromRemote() error {
	httpClient := cart.ApiClient.HttpClient

	resp, err := httpClient.R().SetResult(cart.Cart).Get(cart.Route)
	if err != nil {
		return fmt.Errorf("error while getting detailed cart object from magento2-api: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	return nil
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

		if err != nil {
			return fmt.Errorf("received error while adding item '%v' to cart: '%v'", item, err)
		} else if resp.StatusCode() >= 400 {
			return fmt.Errorf("unexpected statuscode '%v' while adding item '%v' to cart - response: '%v' - body: '%+v'", resp.StatusCode(), item, resp, resp.Request.Body)
		}
	}

	return cart.updateCartFromRemote()
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

	shippingCosts := &[]Carrier{}

	resp, err := httpClient.R().SetBody(*payLoad).SetResult(shippingCosts).Post(endpoint)
	if err != nil {
		return *shippingCosts, fmt.Errorf("received erro while estimating shipping costs: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return *shippingCosts, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	shippingCosts = resp.Result().(*[]Carrier)

	if len(*shippingCosts) == 0 {
		return *shippingCosts, fmt.Errorf("received no suitable shipping - response: '%v'", resp)
	}

	return *shippingCosts, nil
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
	if err != nil {
		return fmt.Errorf("received error while adding shipping information: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	return cart.updateCartFromRemote()
}

func (cart *MCart) EstimatePaymentMethods() ([]PaymentMethod, error) {
	endpoint := cart.Route + cartPaymentMethods
	httpClient := cart.ApiClient.HttpClient

	paymentMethods := &[]PaymentMethod{}

	resp, err := httpClient.R().SetResult(paymentMethods).Get(endpoint)
	if err != nil {
		return *paymentMethods, fmt.Errorf("received error while estimating payment methods costs: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return *paymentMethods, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	paymentMethods = resp.Result().(*[]PaymentMethod)

	if len(*paymentMethods) == 0 {
		return *paymentMethods, fmt.Errorf("received no suitable payment method - response: '%v'", resp)
	}

	return *paymentMethods, nil
}

func (cart *MCart) CreateOrder(paymentMethod PaymentMethod) (*orders.Order, error) {
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
	if err != nil {
		return nil, fmt.Errorf("received error while creating order: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	orderIDString := utils.MayTrimSurroundingQuotes(resp.String())
	orderIDInt, err := strconv.Atoi(orderIDString)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while extracting orderID: '%v'", err)
	}

	return &orders.Order{
		ID:        orderIDInt,
		ApiClient: cart.ApiClient,
		Route:     orders.Orders + "/" + orderIDString,
	}, nil
}

func (cart *MCart) DeleteItem(itemID int) error {
	endpoint := cart.Route + cartItems + "/" + strconv.Itoa(itemID)
	httpClient := cart.ApiClient.HttpClient

	resp, err := httpClient.R().Delete(endpoint)
	if err != nil {
		return fmt.Errorf("received error while creating order: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	return nil
}

func (cart *MCart) DeleteAllItems() error {
	err := cart.updateCartFromRemote()
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
