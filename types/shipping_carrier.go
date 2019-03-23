package types

type Carrier struct {
	CarrierCode  string `json:"carrier_code"`
	MethodCode   string `json:"method_code"`
	CarrierTitle string `json:"carrier_title"`
	MethodTitle  string `json:"method_title"`
	Amount       int    `json:"amount"`
	BaseAmount   int    `json:"base_amount"`
	Available    bool   `json:"available"`
	ErrorMessage string `json:"error_message,omitempty"`
	PriceExclTax int    `json:"price_excl_tax"`
	PriceInclTax int    `json:"price_incl_tax"`
}
