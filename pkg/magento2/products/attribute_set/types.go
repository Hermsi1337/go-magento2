package attribute_set

type AttributeSet struct {
	AttributeSetID      int                    `json:"attribute_set_id,omitempty"`
	AttributeSetName    string                 `json:"attribute_set_name"`
	SortOrder           int                    `json:"sort_order"`
	EntityTypeID        int                    `json:"entity_type_id,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
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
