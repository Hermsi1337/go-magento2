package magento2

import (
	"fmt"
	"github.com/hermsi1337/go-magento2/types"
)

type Order struct {
	ID        int
	Route     string
	ApiClient *ApiClient
}

func (order *Order) AddComment(comment string) error {
	endpoint := order.Route + orderComments
	httpClient := order.ApiClient.HttpClient

	payLoad := types.StatusHistory{
		Comment:  comment,
		ParentID: order.ID,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	if err != nil {
		return fmt.Errorf("received error while creating order: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v'", resp.StatusCode(), resp)
	}

	return nil
}
