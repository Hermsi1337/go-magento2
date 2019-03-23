package types

type AddressInformation struct {
	ShippingAddress      Address                `json:"shipping_address"`
	BillingAddress       Address                `json:"billing_address"`
	ShippingMethodCode   string                 `json:"shipping_method_code"`
	ShippingCarrierCodes string                 `json:"shipping_carrier_code"`
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	CustomAttributes     []CustomAttributes     `json:"custom_attributes,omitempty"`
}
