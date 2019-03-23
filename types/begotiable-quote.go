package types

type NegotiableQuote struct {
	QuoteID                  int                    `json:"quote_id,omitempty"`
	IsRegularQuote           bool                   `json:"is_regular_quote,omitempty"`
	Status                   string                 `json:"status,omitempty"`
	NegotiatedPriceType      int                    `json:"negotiated_price_type,omitempty"`
	NegotiatedPriceValue     int                    `json:"negotiated_price_value,omitempty"`
	ShippingPrice            int                    `json:"shipping_price,omitempty"`
	QuoteName                string                 `json:"quote_name,omitempty"`
	ExpirationPeriod         string                 `json:"expiration_period,omitempty"`
	EmailNotificationStatus  int                    `json:"email_notification_status,omitempty"`
	HasUnconfirmedChanges    bool                   `json:"has_unconfirmed_changes,omitempty"`
	IsShippingTaxChanged     bool                   `json:"is_shipping_tax_changed,omitempty"`
	IsCustomerPriceChanged   bool                   `json:"is_customer_price_changed,omitempty"`
	Notifications            int                    `json:"notifications,omitempty"`
	AppliedRuleIds           string                 `json:"applied_rule_ids,omitempty"`
	IsAddressDraft           bool                   `json:"is_address_draft,omitempty"`
	DeletedSku               string                 `json:"deleted_sku,omitempty"`
	CreatorID                int                    `json:"creator_id,omitempty"`
	CreatorType              int                    `json:"creator_type,omitempty"`
	OriginalTotalPrice       int                    `json:"original_total_price,omitempty"`
	BaseOriginalTotalPrice   int                    `json:"base_original_total_price,omitempty"`
	NegotiatedTotalPrice     int                    `json:"negotiated_total_price,omitempty"`
	BaseNegotiatedTotalPrice int                    `json:"base_negotiated_total_price,omitempty"`
	ExtensionAttributes      map[string]interface{} `json:"extension_attributes,omitempty"`
}
