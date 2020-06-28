package magento2

type ConfigurableProductOption struct {
	ID                  int         `json:"id,omitempty"`
	AttributeID         string      `json:"attribute_id"`
	Label               string      `json:"label"`
	Position            int         `json:"position"`
	IsUseDefault        bool        `json:"is_use_default"`
	Values              []Value     `json:"values"`
	ExtensionAttributes interface{} `json:"extension_attributes,omitempty"`
	ProductID           int         `json:"product_id,omitempty"`
}

type Value struct {
	ValueIndex          int         `json:"value_index,omitempty"`
	ExtensionAttributes interface{} `json:"extension_attributes,omitempty"`
}

type createConfigurableProductByOptionPayload struct {
	Option ConfigurableProductOption `json:"option"`
}

type addChildSKUPayload struct {
	Sku string `json:"childSku"`
}
