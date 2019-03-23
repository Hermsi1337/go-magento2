package types

type ExtensionAttributes struct {
	FileInfo                FileInfo                  `json:"file_info,omitempty"`
	ShippingAssignments     []ShippingAssignments     `json:"shipping_assignments,omitempty"`
	CustomOptions           []CustomOptions           `json:"custom_options,omitempty"`
	BundleOptions           []BundleOptions           `json:"bundle_options,omitempty"`
	DownloadableOption      DownloadableOption        `json:"downloadable_option,omitempty"`
	GiftcardItemOption      GiftcardItemOption        `json:"giftcard_item_option,omitempty"`
	ConfigurableItemOptions []ConfigurableItemOptions `json:"configurable_item_options,omitempty"`
	NegotiableQuoteItem     NegotiableQuoteItem       `json:"negotiable_quote_item,omitempty"`
	CompanyAttributes       CompanyAttributes         `json:"company_attributes,omitempty"`
	IsSubscribed            bool                      `json:"is_subscribed,omitempty"`
	GiftRegistryID          int                       `json:"gift_registry_id,omitempty"`
	NegotiableQuote         NegotiableQuote           `json:"negotiable_quote,omitempty"`
}
