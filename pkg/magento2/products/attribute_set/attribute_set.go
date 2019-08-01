package attribute_set

import (
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"strconv"
)

type MAttributeSet struct {
	Route        string
	AttributeSet *AttributeSet
	ApiClient    *api.Client
}

func CreateAttributeSet(a AttributeSet, skeletonID int, apiClient *api.Client) (*MAttributeSet, error) {
	mAttributeSet := &MAttributeSet{
		AttributeSet: &AttributeSet{},
		ApiClient: apiClient,
	}
	endpoint := productsAttributeSet
	httpClient := apiClient.HttpClient

	payLoad := createAttributeSetPayload{
		AttributeSet: a,
		SkeletonID:   skeletonID,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mAttributeSet.AttributeSet).Post(endpoint)
	mAttributeSet.Route = productsAttributeSet + "/" + strconv.Itoa(mAttributeSet.AttributeSet.AttributeSetID)

	return mAttributeSet, utils.MayReturnErrorForHTTPResponse(err, resp, "create attribute-set")
}

func (mas *MAttributeSet) AssignAttribute(attributeGroupID, sortOrder int, attributeCode string) error {
	endpoint := productsAttributeSetAttributes
	httpClient := mas.ApiClient.HttpClient

	payLoad := assignAttributePayload{
		AttributeSetID: mas.AttributeSet.AttributeSetID,
		AttributeSetGroupID: attributeGroupID,
		AttributeCode: attributeCode,
		SortOrder: sortOrder,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "assign attribute to attribute-set")
}