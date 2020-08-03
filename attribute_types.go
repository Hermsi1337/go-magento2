package magento2

type createAttributePayload struct {
	Attribute Attribute `json:"attribute"`
}

type addOptionPayload struct {
	Option Option `json:"option"`
}

type ExtensionAttributes struct {
	IsPagebuilderEnabled bool `json:"is_pagebuilder_enabled,omitempty"`
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

type Attribute struct {
	ExtensionAttributes       ExtensionAttributes `json:"extension_attributes,omitempty"`
	IsWysiwygEnabled          bool                `json:"is_wysiwyg_enabled,omitempty"`
	IsHTMLAllowedOnFront      bool                `json:"is_html_allowed_on_front,omitempty"`
	UsedForSortBy             bool                `json:"used_for_sort_by,omitempty"`
	IsFilterable              bool                `json:"is_filterable,omitempty"`
	IsFilterableInSearch      bool                `json:"is_filterable_in_search,omitempty"`
	IsUsedInGrid              bool                `json:"is_used_in_grid,omitempty"`
	IsVisibleInGrid           bool                `json:"is_visible_in_grid,omitempty"`
	IsFilterableInGrid        bool                `json:"is_filterable_in_grid,omitempty"`
	Position                  int                 `json:"position,omitempty"`
	ApplyTo                   []string            `json:"apply_to,omitempty"`
	IsSearchable              string              `json:"is_searchable,omitempty"`
	IsVisibleInAdvancedSearch string              `json:"is_visible_in_advanced_search,omitempty"`
	IsComparable              string              `json:"is_comparable,omitempty"`
	IsUsedForPromoRules       string              `json:"is_used_for_promo_rules,omitempty"`
	IsVisibleOnFront          string              `json:"is_visible_on_front,omitempty"`
	UsedInProductListing      string              `json:"used_in_product_listing,omitempty"`
	IsVisible                 bool                `json:"is_visible,omitempty"`
	Scope                     string              `json:"scope,omitempty"`
	AttributeID               int                 `json:"attribute_id,omitempty"`
	AttributeCode             string              `json:"attribute_code"`
	FrontendInput             string              `json:"frontend_input"`
	EntityTypeID              string              `json:"entity_type_id,omitempty"`
	IsRequired                bool                `json:"is_required,omitempty"`
	Options                   []Option            `json:"options,omitempty"`
	IsUserDefined             bool                `json:"is_user_defined,omitempty"`
	DefaultFrontendLabel      string              `json:"default_frontend_label"`
	FrontendLabels            []FrontendLabels    `json:"frontend_labels,omitempty"`
	Note                      string              `json:"note,omitempty"`
	BackendType               string              `json:"backend_type,omitempty"`
	BackendModel              string              `json:"backend_model,omitempty"`
	SourceModel               string              `json:"source_model,omitempty"`
	DefaultValue              string              `json:"default_value,omitempty"`
	IsUnique                  string              `json:"is_unique,omitempty"`
	FrontendClass             string              `json:"frontend_class,omitempty"`
	ValidationRules           []ValidationRules   `json:"validation_rules,omitempty"`
	CustomAttributes          []CustomAttributes  `json:"custom_attributes,omitempty"`
}
