package types

type CustomOptions struct {
	OptionID            string                 `json:"option_id,omitempty"`
	OptionValue         string                 `json:"option_value,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
