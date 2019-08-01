package attribute

import (
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
)

type MAttribute struct {
	Route     string
	Attribute *Attribute
	ApiClient *api.Client
}

func CreateAttribute(a Attribute, apiClient *api.Client) (*MAttribute, error) {
	mAttribute := &MAttribute{
		Attribute: &Attribute{},
		ApiClient: apiClient,
	}
	endpoint := productsAttribute
	httpClient := apiClient.HttpClient

	payLoad := createAttributePayload{
		Attribute: a,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mAttribute.Attribute).Post(endpoint)
	mAttribute.Route = productsAttribute + "/" + mAttribute.Attribute.AttributeCode

	return mAttribute, utils.MayReturnErrorForHTTPResponse(err, resp, "create attribute")
}

func (mas *MAttribute) AddOption(option Option) error {
	endpoint := mas.Route + "/" + productsAttributeOptions
	httpClient := mas.ApiClient.HttpClient

	payLoad := addOptionPayload{
		Option: option,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "assign attribute to attribute-set")
}