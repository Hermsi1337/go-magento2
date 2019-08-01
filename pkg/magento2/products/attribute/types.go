package attribute

type createAttributePayload struct {
	Attribute Attribute `json:"attribute"`
}

type addOptionPayload struct {
	Option Option `json:"option"`
}

type ExtensionAttributes struct {
	IsPagebuilderEnabled bool `json:"is_pagebuilder_enabled"`
}

type StoreLabels struct {
	StoreID int    `json:"store_id"`
	Label   string `json:"label"`
}

type Option struct {
	Label       string        `json:"label"`
	Value       string        `json:"value"`
	SortOrder   int           `json:"sort_order,omitempty"`
	IsDefault   bool          `json:"is_default,omitempty"`
	StoreLabels []StoreLabels `json:"store_labels,omitempty"`
}

type FrontendLabels struct {
	StoreID int    `json:"store_id"`
	Label   string `json:"label"`
}

type ValidationRules struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CustomAttributes struct {
	AttributeCode string `json:"attribute_code"`
	Value         string `json:"value"`
}

type Attribute struct {
	ExtensionAttributes       ExtensionAttributes `json:"extension_attributes"`
	IsWysiwygEnabled          bool                `json:"is_wysiwyg_enabled"`
	IsHTMLAllowedOnFront      bool                `json:"is_html_allowed_on_front"`
	UsedForSortBy             bool                `json:"used_for_sort_by"`
	IsFilterable              bool                `json:"is_filterable"`
	IsFilterableInSearch      bool                `json:"is_filterable_in_search"`
	IsUsedInGrid              bool                `json:"is_used_in_grid"`
	IsVisibleInGrid           bool                `json:"is_visible_in_grid"`
	IsFilterableInGrid        bool                `json:"is_filterable_in_grid"`
	Position                  int                 `json:"position"`
	ApplyTo                   []string            `json:"apply_to"`
	IsSearchable              string              `json:"is_searchable"`
	IsVisibleInAdvancedSearch string              `json:"is_visible_in_advanced_search"`
	IsComparable              string              `json:"is_comparable"`
	IsUsedForPromoRules       string              `json:"is_used_for_promo_rules"`
	IsVisibleOnFront          string              `json:"is_visible_on_front"`
	UsedInProductListing      string              `json:"used_in_product_listing"`
	IsVisible                 bool                `json:"is_visible"`
	Scope                     string              `json:"scope"`
	AttributeID               int                 `json:"attribute_id"`
	AttributeCode             string              `json:"attribute_code"`
	FrontendInput             string              `json:"frontend_input"`
	EntityTypeID              string              `json:"entity_type_id"`
	IsRequired                bool                `json:"is_required"`
	Options                   []Option            `json:"options"`
	IsUserDefined             bool                `json:"is_user_defined"`
	DefaultFrontendLabel      string              `json:"default_frontend_label"`
	FrontendLabels            []FrontendLabels    `json:"frontend_labels"`
	Note                      string              `json:"note"`
	BackendType               string              `json:"backend_type"`
	BackendModel              string              `json:"backend_model"`
	SourceModel               string              `json:"source_model"`
	DefaultValue              string              `json:"default_value"`
	IsUnique                  string              `json:"is_unique"`
	FrontendClass             string              `json:"frontend_class"`
	ValidationRules           []ValidationRules   `json:"validation_rules"`
	CustomAttributes          []CustomAttributes  `json:"custom_attributes"`
}
