package categories

type Category struct {
	ID                  int                `json:"id,omitempty"`
	ParentID            int                `json:"parent_id,omitempty"`
	Name                string             `json:"name,omitempty"`
	IsActive            bool               `json:"is_active,omitempty"`
	Position            int                `json:"position,omitempty"`
	Level               int                `json:"level,omitempty"`
	Children            string             `json:"children,omitempty"`
	CreatedAt           string             `json:"created_at,omitempty"`
	UpdatedAt           string             `json:"updated_at,omitempty"`
	Path                string             `json:"path,omitempty"`
	AvailableSortBy     []string           `json:"available_sort_by,omitempty"`
	IncludeInMenu       bool               `json:"include_in_menu,omitempty"`
	ExtensionAttributes interface{}        `json:"extension_attributes,omitempty"`
	CustomAttributes    []CustomAttributes `json:"custom_attributes,omitempty"`
}

type CustomAttributes struct {
	AttributeCode string `json:"attribute_code"`
	Value         string `json:"value"`
}

type createCategoryPayload struct {
	Category Category `json:"category"`
}

type categorySearchQueryResponse struct {
	Categories     []Category `json:"items"`
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
