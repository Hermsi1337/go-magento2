package types

type Item struct {
	ItemID              int                    `json:"item_id,omitempty"`
	Sku                 string                 `json:"sku"`
	Qty                 int                    `json:"qty"`
	Name                string                 `json:"name,omitempty"`
	Price               int                    `json:"price,omitempty"`
	ProductType         string                 `json:"product_type,omitempty"`
	QuoteID             string                 `json:"quote_id"`
	ProductOption       ProductOption          `json:"product_option,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
