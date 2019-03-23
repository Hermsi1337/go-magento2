package magento2

import (
	"fmt"
)

type Cart struct {
	QuoteID  string
	Detailed DetailedCart
}

func (cart *Cart) GetDetails(apiClient *ApiClient) (DetailedCart, error) {
	var detailedCart = &DetailedCart{}
	httpClient := apiClient.httpClient
	resp, err := httpClient.R().SetResult(&DetailedCart{}).Get(guestCarts + "/" + cart.QuoteID)
	if err != nil {
		return *detailedCart, err
	} else if resp.StatusCode() >= 400 {
		return *detailedCart, fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}
	detailedCart = resp.Result().(*DetailedCart)

	return *detailedCart, err
}

func (cart *Cart) UpdateSelf(apiClient *ApiClient) error {
	newDetails, err := cart.GetDetails(apiClient)
	if err != nil {
		return err
	}

	cart.Detailed = newDetails
	return nil
}

func (cart *Cart) AddItems(apiClient *ApiClient, items []Item) error {
	endpoint := guestCarts + "/" + cart.QuoteID + guestCartsItems
	httpClient := apiClient.httpClient

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
			return err
		} else if resp.StatusCode() >= 400 {
			return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
		}
	}

	return nil
}

func (cart *Cart) AddShippingInformation(apiClient *ApiClient, addrInfo AddressInformation) error {
	endpoint := guestCarts + "/" + cart.QuoteID + guestCartShippingInformation
	httpClient := apiClient.httpClient

	type PayLoad struct {
		AddressInformation AddressInformation `json:"addressInformation"`
	}

	payLoad := &PayLoad{
		AddressInformation: addrInfo,
	}

	resp, err := httpClient.R().SetBody(*payLoad).Post(endpoint)
	if err != nil {
		return err
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	return nil
}
