package types

type Shipping struct {
	Address             Address                `json:"address,omitempty"`
	Method              string                 `json:"method,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
