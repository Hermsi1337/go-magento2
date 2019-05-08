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

func (order *Order) AddComment(statusHistory types.StatusHistory) error {
	endpoint := order.Route + orderComments
	httpClient := order.ApiClient.HttpClient

	type PayLoad struct {
		StatusHistory types.StatusHistory `json:"statusHistory"`
	}

	payLoad := &PayLoad{
		StatusHistory: statusHistory,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	if err != nil {
		return fmt.Errorf("received error while adding comment to order: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v' - body: '%+v'", resp.StatusCode(), resp, resp.Request.Body)
	}

	return nil
}
