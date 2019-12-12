package orders

import (
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
)

type MOrder struct {
	Route     string
	Order     *Order
	ApiClient *api.Client
}

func (mo *MOrder) UpdateOrderFromRemote() error {
	httpClient := mo.ApiClient.HttpClient

	resp, err := httpClient.R().SetResult(mo.Order).Get(mo.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get detailed order object from magento2-api")
}

func (mo *MOrder) AddComment(comment StatusHistory) error {
	endpoint := mo.Route + "/" + OrderComments
	httpClient := mo.ApiClient.HttpClient

	type PayLoad struct {
		StatusHistory StatusHistory `json:"statusHistory"`
	}

	payLoad := &PayLoad{
		StatusHistory: comment,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "add comment to order")
}
