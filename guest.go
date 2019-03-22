package magento2

import (
	"gopkg.in/resty.v1"
)

type GuestApiClient struct {
	httpClient *resty.Client
}

type CartID string

func NewGuestApiClient(scheme string, hostName string) *GuestApiClient {
	client := buildBasicApiClient(scheme, hostName)

	return &GuestApiClient{
		httpClient: client,
	}
}

func (client *GuestApiClient) CreateEmptyCartID() (CartID, error) {
	resp, err := client.httpClient.R().Post(guestCarts)
	if err != nil {
		return "", err
	}
	return CartID(resp.String()), err
}

func (client *GuestApiClient) GetCartByID(id CartID) (DetailedCart, error) {
	var detailedCart *DetailedCart
	resp, err := client.httpClient.R().SetResult(&DetailedCart{}).Get(guestCarts + "/" + string(id))
	if err != nil {
		return *detailedCart, err
	}
	detailedCart = resp.Result().(*DetailedCart)

	return *detailedCart, err
}
