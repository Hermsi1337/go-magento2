package types

type ConfigurableItemOptions struct {
	OptionID            string                 `json:"option_id,omitempty"`
	OptionValue         int                    `json:"option_value,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
