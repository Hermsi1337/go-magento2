package attribute_set

import (
	"github.com/hermsi1337/go-magento2/internal/utils"
	"github.com/hermsi1337/go-magento2/pkg/magento2"
	"github.com/hermsi1337/go-magento2/pkg/magento2/api"
	"github.com/hermsi1337/go-magento2/pkg/magento2/products/attribute"
	"strconv"
)

type MAttributeSet struct {
	Route                  string
	AttributeSet           *AttributeSet
	AttributeSetGroups     []AttributeSetGroup
	AttributeSetAttributes *[]attribute.Attribute
	ApiClient              *api.Client
}

func CreateAttributeSet(a AttributeSet, skeletonID int, apiClient *api.Client) (*MAttributeSet, error) {
	mAttributeSet := &MAttributeSet{
		AttributeSet:           &AttributeSet{},
		AttributeSetAttributes: &[]attribute.Attribute{},
		ApiClient:              apiClient,
	}
	endpoint := productsAttributeSet
	httpClient := apiClient.HttpClient

	payLoad := createAttributeSetPayload{
		AttributeSet: a,
		SkeletonID:   skeletonID,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mAttributeSet.AttributeSet).Post(endpoint)
	mAttributeSet.Route = productsAttributeSet + "/" + strconv.Itoa(mAttributeSet.AttributeSet.AttributeSetID)

	err = utils.MayReturnErrorForHTTPResponse(err, resp, "create attribute-set")
	if err != nil {
		return mAttributeSet, err
	}

	err = mAttributeSet.UpdateAttributeSetFromRemote()

	return mAttributeSet, err
}

func GetAttributeSetByName(name string, apiClient *api.Client) (*MAttributeSet, error) {
	mAttributeSet := &MAttributeSet{
		AttributeSet:           &AttributeSet{},
		AttributeSetAttributes: &[]attribute.Attribute{},
		ApiClient:              apiClient,
	}
	searchQuery := utils.BuildSearchQuery("attribute_set_name", name, "in")
	endpoint := productsAttributeSetList + "?" + searchQuery
	httpClient := apiClient.HttpClient

	response := &attributeSetSearchQueryResponse{}

	resp, err := httpClient.R().SetResult(response).Get(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "get attribute-set by name from remote")
	if err != nil {
		return nil, err
	}

	if len(response.AttributeSets) <= 0 {
		return nil, magento2.ErrNotFound
	}

	mAttributeSet.AttributeSet = &response.AttributeSets[0]
	mAttributeSet.Route = productsAttributeSet + "/" + strconv.Itoa(mAttributeSet.AttributeSet.AttributeSetID)
	err = utils.MayReturnErrorForHTTPResponse(mAttributeSet.UpdateAttributeSetFromRemote(), resp, "get detailed attribute-set by name from remote")

	return mAttributeSet, err
}

func (mas *MAttributeSet) UpdateAttributeSetOnRemote() error {
	resp, err := mas.ApiClient.HttpClient.R().SetResult(mas.AttributeSet).SetBody(mas.AttributeSet).Put(mas.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "update remote attribute-set from local")
}

func (mas *MAttributeSet) UpdateAttributeSetFromRemote() error {
	err := mas.updateAttributeSet()
	if err != nil {
		return err
	}

	err = mas.updateGroups()
	if err != nil {
		return err
	}

	return mas.updateAttributes()
}

func (mas *MAttributeSet) updateAttributeSet() error {
	resp, err := mas.ApiClient.HttpClient.R().SetResult(mas.AttributeSet).Get(mas.Route)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get details for attribute-set from remote")
}

func (mas *MAttributeSet) updateAttributes() error {
	resp, err := mas.ApiClient.HttpClient.R().SetResult(mas.AttributeSetAttributes).Get(mas.Route + "/" + productsAttributeSetAttributesRelative)
	return utils.MayReturnErrorForHTTPResponse(err, resp, "get details for attribute-set from remote")
}

func (mas *MAttributeSet) updateGroups() error {
	searchQuery := utils.BuildSearchQuery("attribute_set_id", strconv.Itoa(mas.AttributeSet.AttributeSetID), "in")
	endpoint := productsAttributeSetGroupsList + "?" + searchQuery

	response := &groupSearchQueryResponse{}

	resp, err := mas.ApiClient.HttpClient.R().SetResult(response).Get(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "get groups for attribute-set from remote")
	if err != nil {
		return err
	}

	mas.AttributeSetGroups = response.Groups

	return nil
}

func (mas *MAttributeSet) AssignAttribute(attributeGroupID, sortOrder int, attributeCode string) error {
	endpoint := productsAttributeSetAttributes
	httpClient := mas.ApiClient.HttpClient

	payLoad := assignAttributePayload{
		AttributeSetID:      mas.AttributeSet.AttributeSetID,
		AttributeSetGroupID: attributeGroupID,
		AttributeCode:       attributeCode,
		SortOrder:           sortOrder,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "assign attribute to attribute-set")
	if err != nil {
		return err
	}

	return mas.UpdateAttributeSetFromRemote()
}

func (mas *MAttributeSet) CreateGroup(groupName string) error {
	endpoint := productsAttributeSetGroups
	httpClient := mas.ApiClient.HttpClient

	payLoad := createGroupPayload{
		Group: AttributeSetGroup{
			AttributeGroupName: groupName,
			AttributeSetID:     mas.AttributeSet.AttributeSetID,
		},
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	err = utils.MayReturnErrorForHTTPResponse(err, resp, "create group on attribute-set")
	if err != nil {
		return err
	}

	return mas.UpdateAttributeSetFromRemote()
}
