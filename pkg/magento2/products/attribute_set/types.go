package attribute_set

import "errors"

var ErrNotFound = errors.New("no document found")

type AttributeSet struct {
	AttributeSetID      int         `json:"attribute_set_id,omitempty"`
	AttributeSetName    string      `json:"attribute_set_name"`
	SortOrder           int         `json:"sort_order"`
	EntityTypeID        int         `json:"entity_type_id,omitempty"`
	ExtensionAttributes interface{} `json:"extension_attributes,omitempty"`
}

type AttributeSetGroups struct {
	AttributeGroupID    string `json:"attribute_group_id"`
	AttributeGroupName  string `json:"attribute_group_name"`
	AttributeSetID      int    `json:"attribute_set_id"`
	ExtensionAttributes struct {
		AttributeGroupCode string `json:"attribute_group_code"`
		SortOrder          string `json:"sort_order"`
	} `json:"extension_attributes"`
}

type attributeSetSearchQueryResponse struct {
	AttributeSets  []AttributeSet `json:"items"`
	SearchCriteria struct {
		FilterGroups []struct {
			Filters []struct {
				Field         string `json:"field"`
				Value         string `json:"value"`
				ConditionType string `json:"condition_type"`
			} `json:"filters"`
		} `json:"filter_groups"`
	} `json:"search_criteria"`
}

type groupSearchQueryResponse struct {
	Groups         []AttributeSetGroups `json:"items"`
	SearchCriteria struct {
		FilterGroups []struct {
			Filters []struct {
				Field         string `json:"field"`
				Value         string `json:"value"`
				ConditionType string `json:"condition_type"`
			} `json:"filters"`
		} `json:"filter_groups"`
	} `json:"search_criteria"`
}

type createAttributeSetPayload struct {
	AttributeSet AttributeSet `json:"attributeSet"`
	SkeletonID   int          `json:"skeletonId"`
}

type assignAttributePayload struct {
	AttributeSetID      int    `json:"attributeSetId"`
	AttributeSetGroupID int    `json:"attributeGroupId"`
	AttributeCode       string `json:"attributeCode"`
	SortOrder           int    `json:"sortOrder"`
}
