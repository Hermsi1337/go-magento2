package types

type Currency struct {
	GlobalCurrencyCode  string                 `json:"global_currency_code,omitempty"`
	BaseCurrencyCode    string                 `json:"base_currency_code,omitempty"`
	StoreCurrencyCode   string                 `json:"store_currency_code,omitempty"`
	QuoteCurrencyCode   string                 `json:"quote_currency_code,omitempty"`
	StoreToBaseRate     int                    `json:"store_to_base_rate,omitempty"`
	StoreToQuoteRate    int                    `json:"store_to_quote_rate,omitempty"`
	BaseToGlobalRate    int                    `json:"base_to_global_rate,omitempty"`
	BaseToQuoteRate     int                    `json:"base_to_quote_rate,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
