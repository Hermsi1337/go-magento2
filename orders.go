package magento2

type MOrder struct {
	Route     string
	Order     *Order
	APIClient *Client
}

func GetOrderByIncrementID(id string, apiClient *Client) (*MOrder, error) {
	mOrder := &MOrder{
		Route:     "",
		Order:     &Order{},
		APIClient: apiClient,
	}

	// ?searchCriteria[filter_groups][2][filters][0][field]=increment_id
	// &searchCriteria[filter_groups][2][filters][0][value]=INCREMENT_ID_HERE
	// &searchCriteria[filter_groups][2][filters][0][condition_type]=eq
	// &fields=items[entity_id]
	searchCriteria := []SearchQueryCriteria{
		{
			Fields: []FilterFields{
				{
					Field: Filter{
						FilterGroups: 2,
						Filters:      0,
						FilterFor:    "increment_id",
					},
					Value: Filter{
						FilterGroups: 2,
						Filters:      0,
						FilterFor:    id,
					},
					ConditionType: Filter{
						FilterGroups: 2,
						Filters:      0,
						FilterFor:    "eq",
					},
				},
			},
		},
	}

	additionalQuery := Fields{
		Key:   "fields",
		Value: "items[entity_id]",
	}

	searchQuery := BuildFlexibleSearchQuery(searchCriteria, additionalQuery)

	type searchResponse struct {
		Items []struct {
			EntityID int `json:"entity_id"`
		}
	}

	response := &searchResponse{
		Items: []struct {
			EntityID int `json:"entity_id"`
		}{},
	}

	endpoint := Orders + "?" + searchQuery

	err := apiClient.GetRouteAndDecode(endpoint, response, "get order by increment_id from remote")
	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, ErrNotFound
	}

	mOrder.Order.EntityID = response.Items[0].EntityID
	err = mOrder.UpdateFromRemote()

	return mOrder, err
}

func (mo *MOrder) UpdateEntity(order *Order) error {
	type updateOrderEntityPayload struct {
		Entity Order `json:"entity"`
	}

	order.EntityID = mo.Order.EntityID

	payLoad := updateOrderEntityPayload{
		Entity: *order,
	}

	resp, err := mo.APIClient.HTTPClient.R().SetResult(mo.Order).SetBody(payLoad).Post(Orders)
	return mayReturnErrorForHTTPResponse(err, resp, "update order entity on remote")
}

func (mo *MOrder) UpdateFromRemote() error {
	return mo.APIClient.GetRouteAndDecode(mo.Route, mo.Order, "get detailed order object from magento2-api")
}

func (mo *MOrder) AddComment(comment *StatusHistory) (StatusHistory, error) {
	endpoint := mo.Route + "/" + OrderComments

	type PayLoad struct {
		StatusHistory StatusHistory `json:"statusHistory"`
	}

	payLoad := &PayLoad{
		StatusHistory: *comment,
	}

	response := StatusHistory{}

	err := mo.APIClient.PostRouteAndDecode(endpoint, payLoad, &response, "add comment to order")
	return response, err
}
