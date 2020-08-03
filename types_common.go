package magento2

type Address struct {
	ID                  int                      `json:"id,omitempty"`
	RegionID            int                      `json:"region_id,omitempty"`
	RegionCode          string                   `json:"region_code,omitempty"`
	CountryID           string                   `json:"country_id"`
	Street              []string                 `json:"street"`
	Company             string                   `json:"company,omitempty"`
	Telephone           string                   `json:"telephone"`
	Fax                 string                   `json:"fax,omitempty"`
	Postcode            string                   `json:"postcode"`
	City                string                   `json:"city"`
	Firstname           string                   `json:"firstname"`
	Lastname            string                   `json:"lastname"`
	Middlename          string                   `json:"middlename,omitempty"`
	Prefix              string                   `json:"prefix,omitempty"`
	Suffix              string                   `json:"suffix,omitempty"`
	VatID               string                   `json:"vat_id,omitempty"`
	CustomerID          int                      `json:"customer_id,omitempty"`
	Email               string                   `json:"email"`
	SameAsBilling       int                      `json:"same_as_billing,omitempty"`
	CustomerAddressID   int                      `json:"customer_address_id,omitempty"`
	SaveInAddressBook   int                      `json:"save_in_address_book,omitempty"`
	ExtensionAttributes map[string]interface{}   `json:"extension_attributes,omitempty"`
	CustomAttributes    []map[string]interface{} `json:"custom_attributes,omitempty"`
}

type BillingAddress struct {
	Address
}

type ShippingAddress struct {
	Address
	AddressType       string  `json:"address_type,omitempty"`
	EntityID          int     `json:"entity_id,omitempty"`
	ParentID          float64 `json:"parent_id,omitempty"`
	Region            string  `json:"region,omitempty"`
	VatIsValid        float64 `json:"vat_is_valid,omitempty"`
	VatRequestDate    string  `json:"vat_request_date,omitempty"`
	VatRequestID      string  `json:"vat_request_id,omitempty"`
	VatRequestSuccess float64 `json:"vat_request_success,omitempty"`
}

type Item struct {
	AdditionalData                      string  `json:"additional_data,omitempty"`
	AmountRefunded                      float64 `json:"amount_refunded,omitempty"`
	AppliedRuleIds                      string  `json:"applied_rule_ids,omitempty"`
	BaseAmountRefunded                  float64 `json:"base_amount_refunded,omitempty"`
	BaseCost                            float64 `json:"base_cost,omitempty"`
	BaseDiscountAmount                  float64 `json:"base_discount_amount,omitempty"`
	BaseDiscountInvoiced                float64 `json:"base_discount_invoiced,omitempty"`
	BaseDiscountRefunded                float64 `json:"base_discount_refunded,omitempty"`
	BaseDiscountTaxCompensationAmount   float64 `json:"base_discount_tax_compensation_amount,omitempty"`
	BaseDiscountTaxCompensationInvoiced float64 `json:"base_discount_tax_compensation_invoiced,omitempty"`
	BaseDiscountTaxCompensationRefunded float64 `json:"base_discount_tax_compensation_refunded,omitempty"`
	BaseOriginalPrice                   float64 `json:"base_original_price,omitempty"`
	BasePrice                           float64 `json:"base_price,omitempty"`
	BasePriceInclTax                    float64 `json:"base_price_incl_tax,omitempty"`
	BaseRowInvoiced                     float64 `json:"base_row_invoiced,omitempty"`
	BaseRowTotal                        float64 `json:"base_row_total,omitempty"`
	BaseRowTotalInclTax                 float64 `json:"base_row_total_incl_tax,omitempty"`
	BaseTaxAmount                       float64 `json:"base_tax_amount,omitempty"`
	BaseTaxBeforeDiscount               float64 `json:"base_tax_before_discount,omitempty"`
	BaseTaxInvoiced                     float64 `json:"base_tax_invoiced,omitempty"`
	BaseTaxRefunded                     float64 `json:"base_tax_refunded,omitempty"`
	BaseWeeeTaxAppliedAmount            float64 `json:"base_weee_tax_applied_amount,omitempty"`
	BaseWeeeTaxAppliedRowAmnt           float64 `json:"base_weee_tax_applied_row_amnt,omitempty"`
	BaseWeeeTaxDisposition              float64 `json:"base_weee_tax_disposition,omitempty"`
	BaseWeeeTaxRowDisposition           float64 `json:"base_weee_tax_row_disposition,omitempty"`
	CreatedAt                           string  `json:"created_at,omitempty"`
	Description                         string  `json:"description,omitempty"`
	DiscountAmount                      float64 `json:"discount_amount,omitempty"`
	DiscountInvoiced                    float64 `json:"discount_invoiced,omitempty"`
	DiscountPercent                     float64 `json:"discount_percent,omitempty"`
	DiscountRefunded                    float64 `json:"discount_refunded,omitempty"`
	EventID                             float64 `json:"event_id,omitempty"`
	ExtOrderItemID                      string  `json:"ext_order_item_id,omitempty"`
	FreeShipping                        float64 `json:"free_shipping,omitempty"`
	GwBasePrice                         float64 `json:"gw_base_price,omitempty"`
	GwBasePriceInvoiced                 float64 `json:"gw_base_price_invoiced,omitempty"`
	GwBasePriceRefunded                 float64 `json:"gw_base_price_refunded,omitempty"`
	GwBaseTaxAmount                     float64 `json:"gw_base_tax_amount,omitempty"`
	GwBaseTaxAmountInvoiced             float64 `json:"gw_base_tax_amount_invoiced,omitempty"`
	GwBaseTaxAmountRefunded             float64 `json:"gw_base_tax_amount_refunded,omitempty"`
	GwID                                float64 `json:"gw_id,omitempty"`
	GwPrice                             float64 `json:"gw_price,omitempty"`
	GwPriceInvoiced                     float64 `json:"gw_price_invoiced,omitempty"`
	GwPriceRefunded                     float64 `json:"gw_price_refunded,omitempty"`
	GwTaxAmount                         float64 `json:"gw_tax_amount,omitempty"`
	GwTaxAmountInvoiced                 float64 `json:"gw_tax_amount_invoiced,omitempty"`
	GwTaxAmountRefunded                 float64 `json:"gw_tax_amount_refunded,omitempty"`
	DiscountTaxCompensationAmount       float64 `json:"discount_tax_compensation_amount,omitempty"`
	DiscountTaxCompensationCanceled     float64 `json:"discount_tax_compensation_canceled,omitempty"`
	DiscountTaxCompensationInvoiced     float64 `json:"discount_tax_compensation_invoiced,omitempty"`
	DiscountTaxCompensationRefunded     float64 `json:"discount_tax_compensation_refunded,omitempty"`
	IsQtyDecimal                        float64 `json:"is_qty_decimal,omitempty"`
	IsVirtual                           float64 `json:"is_virtual,omitempty"`
	ItemID                              float64 `json:"item_id,omitempty"`
	LockedDoInvoice                     float64 `json:"locked_do_invoice,omitempty"`
	LockedDoShip                        float64 `json:"locked_do_ship,omitempty"`
	Name                                string  `json:"name,omitempty"`
	NoDiscount                          float64 `json:"no_discount,omitempty"`
	OrderID                             float64 `json:"order_id,omitempty"`
	OriginalPrice                       float64 `json:"original_price,omitempty"`
	ParentItemID                        float64 `json:"parent_item_id,omitempty"`
	Price                               float64 `json:"price,omitempty"`
	PriceInclTax                        float64 `json:"price_incl_tax,omitempty"`
	ProductID                           float64 `json:"product_id,omitempty"`
	ProductType                         string  `json:"product_type,omitempty"`
	QtyBackordered                      float64 `json:"qty_backordered,omitempty"`
	QtyCanceled                         float64 `json:"qty_canceled,omitempty"`
	QtyInvoiced                         float64 `json:"qty_invoiced,omitempty"`
	QtyOrdered                          float64 `json:"qty_ordered,omitempty"`
	QtyRefunded                         float64 `json:"qty_refunded,omitempty"`
	QtyReturned                         float64 `json:"qty_returned,omitempty"`
	QtyShipped                          float64 `json:"qty_shipped,omitempty"`
	QuoteItemID                         float64 `json:"quote_item_id,omitempty"`
	RowInvoiced                         float64 `json:"row_invoiced,omitempty"`
	RowTotal                            float64 `json:"row_total,omitempty"`
	RowTotalInclTax                     float64 `json:"row_total_incl_tax,omitempty"`
	RowWeight                           float64 `json:"row_weight,omitempty"`
	Sku                                 string  `json:"sku,omitempty"`
	StoreID                             float64 `json:"store_id,omitempty"`
	TaxAmount                           float64 `json:"tax_amount,omitempty"`
	TaxBeforeDiscount                   float64 `json:"tax_before_discount,omitempty"`
	TaxCanceled                         float64 `json:"tax_canceled,omitempty"`
	TaxInvoiced                         float64 `json:"tax_invoiced,omitempty"`
	TaxPercent                          float64 `json:"tax_percent,omitempty"`
	TaxRefunded                         float64 `json:"tax_refunded,omitempty"`
	UpdatedAt                           string  `json:"updated_at,omitempty"`
	WeeeTaxApplied                      string  `json:"weee_tax_applied,omitempty"`
	WeeeTaxAppliedAmount                float64 `json:"weee_tax_applied_amount,omitempty"`
	WeeeTaxAppliedRowAmount             float64 `json:"weee_tax_applied_row_amount,omitempty"`
	WeeeTaxDisposition                  float64 `json:"weee_tax_disposition,omitempty"`
	WeeeTaxRowDisposition               float64 `json:"weee_tax_row_disposition,omitempty"`
	Weight                              float64 `json:"weight,omitempty"`
	ParentItem                          *struct {
	} `json:"parent_item,omitempty"`
	ProductOption       OrdersProductOption `json:"product_option,omitempty"`
	ExtensionAttributes *struct {
		GiftMessage *struct {
			GiftMessageID       float64 `json:"gift_message_id,omitempty"`
			CustomerID          float64 `json:"customer_id,omitempty"`
			Sender              string  `json:"sender,omitempty"`
			Recipient           string  `json:"recipient,omitempty"`
			Message             string  `json:"message,omitempty"`
			ExtensionAttributes *struct {
				EntityID                   string  `json:"entity_id,omitempty"`
				EntityType                 string  `json:"entity_type,omitempty"`
				WrappingID                 float64 `json:"wrapping_id,omitempty"`
				WrappingAllowGiftReceipt   bool    `json:"wrapping_allow_gift_receipt,omitempty"`
				WrappingAddPrfloat64edCard bool    `json:"wrapping_add_prfloat64ed_card,omitempty"`
			} `json:"extension_attributes,omitempty"`
		} `json:"gift_message,omitempty"`
		GwID                    string   `json:"gw_id,omitempty"`
		GwBasePrice             string   `json:"gw_base_price,omitempty"`
		GwPrice                 string   `json:"gw_price,omitempty"`
		GwBaseTaxAmount         string   `json:"gw_base_tax_amount,omitempty"`
		GwTaxAmount             string   `json:"gw_tax_amount,omitempty"`
		GwBasePriceInvoiced     string   `json:"gw_base_price_invoiced,omitempty"`
		GwPriceInvoiced         string   `json:"gw_price_invoiced,omitempty"`
		GwBaseTaxAmountInvoiced string   `json:"gw_base_tax_amount_invoiced,omitempty"`
		GwTaxAmountInvoiced     string   `json:"gw_tax_amount_invoiced,omitempty"`
		GwBasePriceRefunded     string   `json:"gw_base_price_refunded,omitempty"`
		GwPriceRefunded         string   `json:"gw_price_refunded,omitempty"`
		GwBaseTaxAmountRefunded string   `json:"gw_base_tax_amount_refunded,omitempty"`
		GwTaxAmountRefunded     string   `json:"gw_tax_amount_refunded,omitempty"`
		VertexTaxCodes          []string `json:"vertex_tax_codes,omitempty"`
		InvoiceTextCodes        []string `json:"invoice_text_codes,omitempty"`
		TaxCodes                []string `json:"tax_codes,omitempty"`
	} `json:"extension_attributes,omitempty"`
}
