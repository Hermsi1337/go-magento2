package types

type BundleOptions struct {
	OptionID            int                    `json:"option_id,omitempty"`
	OptionQty           int                    `json:"option_qty,omitempty"`
	OptionSelections    []int                  `json:"option_selections,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
