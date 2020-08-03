package magento2

import (
	"strconv"
)

type MAttributeSet struct {
	Route                  string
	AttributeSet           *AttributeSet
	AttributeSetGroups     []Group
	AttributeSetAttributes *[]Attribute
	APIClient              *Client
}

func CreateAttributeSet(a AttributeSet, skeletonID int, apiClient *Client) (*MAttributeSet, error) {
	mAttributeSet := &MAttributeSet{
		AttributeSet:           &AttributeSet{},
		AttributeSetAttributes: &[]Attribute{},
		APIClient:              apiClient,
	}
	endpoint := productsAttributeSet
	httpClient := apiClient.HTTPClient

	payLoad := createAttributeSetPayload{
		AttributeSet: a,
		SkeletonID:   skeletonID,
	}

	resp, err := httpClient.R().SetBody(payLoad).SetResult(mAttributeSet.AttributeSet).Post(endpoint)
	mAttributeSet.Route = productsAttributeSet + "/" + strconv.Itoa(mAttributeSet.AttributeSet.AttributeSetID)

	err = mayReturnErrorForHTTPResponse(err, resp, "create attribute-set")
	if err != nil {
		return mAttributeSet, err
	}

	err = mAttributeSet.UpdateAttributeSetFromRemote()

	return mAttributeSet, err
}

func GetAttributeSetByName(name string, apiClient *Client) (*MAttributeSet, error) {
	mAttributeSet := &MAttributeSet{
		AttributeSet:           &AttributeSet{},
		AttributeSetAttributes: &[]Attribute{},
		APIClient:              apiClient,
	}
	searchQuery := BuildSearchQuery("attribute_set_name", name, "in")
	endpoint := productsAttributeSetList + "?" + searchQuery
	httpClient := apiClient.HTTPClient

	response := &attributeSetSearchQueryResponse{}

	resp, err := httpClient.R().SetResult(response).Get(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "get attribute-set by name from remote")
	if err != nil {
		return nil, err
	}

	if len(response.AttributeSets) == 0 {
		return nil, ErrNotFound
	}

	mAttributeSet.AttributeSet = &response.AttributeSets[0]
	mAttributeSet.Route = productsAttributeSet + "/" + strconv.Itoa(mAttributeSet.AttributeSet.AttributeSetID)
	err = mayReturnErrorForHTTPResponse(mAttributeSet.UpdateAttributeSetFromRemote(), resp, "get detailed attribute-set by name from remote")

	return mAttributeSet, err
}

func (mas *MAttributeSet) UpdateAttributeSetOnRemote() error {
	resp, err := mas.APIClient.HTTPClient.R().SetResult(mas.AttributeSet).SetBody(mas.AttributeSet).Put(mas.Route)
	return mayReturnErrorForHTTPResponse(err, resp, "update remote attribute-set from local")
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
	resp, err := mas.APIClient.HTTPClient.R().SetResult(mas.AttributeSet).Get(mas.Route)
	return mayReturnErrorForHTTPResponse(err, resp, "get details for attribute-set from remote")
}

func (mas *MAttributeSet) updateAttributes() error {
	resp, err := mas.APIClient.HTTPClient.R().SetResult(mas.AttributeSetAttributes).Get(mas.Route + "/" + productsAttributeSetAttributesRelative)
	return mayReturnErrorForHTTPResponse(err, resp, "get details for attribute-set from remote")
}

func (mas *MAttributeSet) updateGroups() error {
	searchQuery := BuildSearchQuery("attribute_set_id", strconv.Itoa(mas.AttributeSet.AttributeSetID), "in")
	endpoint := productsAttributeSetGroupsList + "?" + searchQuery

	response := &groupSearchQueryResponse{}

	resp, err := mas.APIClient.HTTPClient.R().SetResult(response).Get(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "get groups for attribute-set from remote")
	if err != nil {
		return err
	}

	mas.AttributeSetGroups = response.Groups

	return nil
}

func (mas *MAttributeSet) AssignAttribute(attributeGroupID, sortOrder int, attributeCode string) error {
	endpoint := productsAttributeSetAttributes
	httpClient := mas.APIClient.HTTPClient

	payLoad := assignAttributePayload{
		AttributeSetID:      mas.AttributeSet.AttributeSetID,
		AttributeSetGroupID: attributeGroupID,
		AttributeCode:       attributeCode,
		SortOrder:           sortOrder,
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "assign attribute to attribute-set")
	if err != nil {
		return err
	}

	return mas.UpdateAttributeSetFromRemote()
}

func (mas *MAttributeSet) CreateGroup(groupName string) error {
	endpoint := productsAttributeSetGroups
	httpClient := mas.APIClient.HTTPClient

	payLoad := createGroupPayload{
		Group: Group{
			AttributeGroupName: groupName,
			AttributeSetID:     mas.AttributeSet.AttributeSetID,
		},
	}

	resp, err := httpClient.R().SetBody(payLoad).Post(endpoint)
	err = mayReturnErrorForHTTPResponse(err, resp, "create group on attribute-set")
	if err != nil {
		return err
	}

	return mas.UpdateAttributeSetFromRemote()
}
