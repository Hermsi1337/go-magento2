package types

type ShippingAssignments struct {
	Shipping            Shipping               `json:"shipping,omitempty"`
	Items               []Item                 `json:"items,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
