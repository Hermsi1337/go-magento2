package comments

import (
	"fmt"
	"github.com/hermsi1337/go-magento2/pkg/magento2/orders"
)

func AddCommentToOrder(comment orders.StatusHistory, order orders.Order) error {
	endpoint := order.Route + OrderComments
	httpClient := order.ApiClient.HttpClient

	type PayLoad struct {
		StatusHistory orders.StatusHistory `json:"statusHistory"`
	}

	payLoad := &PayLoad{
		StatusHistory: comment,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	if err != nil {
		return fmt.Errorf("received error while adding comment to order: '%v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode '%v' - response: '%v' - body: '%+v'", resp.StatusCode(), resp, resp.Request.Body)
	}

	return nil
}
