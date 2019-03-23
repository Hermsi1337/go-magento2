package types

type NegotiableQuoteItem struct {
	ItemID                 int                    `json:"item_id,omitempty"`
	OriginalPrice          int                    `json:"original_price,omitempty"`
	OriginalTaxAmount      int                    `json:"original_tax_amount,omitempty"`
	OriginalDiscountAmount int                    `json:"original_discount_amount,omitempty"`
	ExtensionAttributes    map[string]interface{} `json:"extension_attributes,omitempty"`
}
