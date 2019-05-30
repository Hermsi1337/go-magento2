package magento2

type Address struct {
	ID                  int                    `json:"id,omitempty"`
	RegionID            int                    `json:"region_id,omitempty"`
	RegionCode          string                 `json:"region_code,omitempty"`
	CountryID           string                 `json:"country_id"`
	Street              []string               `json:"street"`
	Company             string                 `json:"company,omitempty"`
	Telephone           string                 `json:"telephone"`
	Fax                 string                 `json:"fax,omitempty"`
	Postcode            string                 `json:"postcode"`
	City                string                 `json:"city"`
	Firstname           string                 `json:"firstname"`
	Lastname            string                 `json:"lastname"`
	Middlename          string                 `json:"middlename,omitempty"`
	Prefix              string                 `json:"prefix,omitempty"`
	Suffix              string                 `json:"suffix,omitempty"`
	VatID               string                 `json:"vat_id,omitempty"`
	CustomerID          int                    `json:"customer_id,omitempty"`
	Email               string                 `json:"email"`
	SameAsBilling       int                    `json:"same_as_billing,omitempty"`
	CustomerAddressID   int                    `json:"customer_address_id,omitempty"`
	SaveInAddressBook   int                    `json:"save_in_address_book,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
	CustomAttributes    []CustomAttributes     `json:"custom_attributes,omitempty"`
}

type AddressInformation struct {
	ShippingAddress      Address                `json:"shipping_address"`
	BillingAddress       Address                `json:"billing_address"`
	ShippingMethodCode   string                 `json:"shipping_method_code"`
	ShippingCarrierCodes string                 `json:"shipping_carrier_code"`
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	CustomAttributes     []CustomAttributes     `json:"custom_attributes,omitempty"`
}

type AuthenticationRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

type BundleOptions struct {
	OptionID            int                    `json:"option_id,omitempty"`
	OptionQty           int                    `json:"option_qty,omitempty"`
	OptionSelections    []int                  `json:"option_selections,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type CompanyAttributes struct {
	CustomerID          int                    `json:"customer_id,omitempty"`
	CompanyID           int                    `json:"company_id,omitempty"`
	JobTitle            string                 `json:"job_title,omitempty"`
	Status              int                    `json:"status,omitempty"`
	Telephone           string                 `json:"telephone,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type ConfigurableItemOptions struct {
	OptionID            string                 `json:"option_id,omitempty"`
	OptionValue         int                    `json:"option_value,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

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

type CustomAttributes struct {
	AttributeCode string `json:"attribute_code,omitempty"`
	Value         string `json:"value,omitempty"`
}

type CustomOptions struct {
	OptionID            string                 `json:"option_id,omitempty"`
	OptionValue         string                 `json:"option_value,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type Customer struct {
	ID                     int                    `json:"id,omitempty"`
	GroupID                int                    `json:"group_id,omitempty"`
	DefaultBilling         string                 `json:"default_billing,omitempty"`
	DefaultShipping        string                 `json:"default_shipping,omitempty"`
	Confirmation           string                 `json:"confirmation,omitempty"`
	CreatedAt              string                 `json:"created_at,omitempty"`
	UpdatedAt              string                 `json:"updated_at,omitempty"`
	CreatedIn              string                 `json:"created_in,omitempty"`
	Dob                    string                 `json:"dob,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Firstname              string                 `json:"firstname,omitempty"`
	Lastname               string                 `json:"lastname,omitempty"`
	Middlename             string                 `json:"middlename,omitempty"`
	Prefix                 string                 `json:"prefix,omitempty"`
	Suffix                 string                 `json:"suffix,omitempty"`
	Gender                 int                    `json:"gender,omitempty"`
	StoreID                int                    `json:"store_id,omitempty"`
	Taxvat                 string                 `json:"taxvat,omitempty"`
	WebsiteID              int                    `json:"website_id,omitempty"`
	Addresses              []Address              `json:"addresses,omitempty"`
	DisableAutoGroupChange int                    `json:"disable_auto_group_change,omitempty"`
	ExtensionAttributes    map[string]interface{} `json:"extension_attributes,omitempty"`
	CustomAttributes       []CustomAttributes     `json:"custom_attributes,omitempty"`
}

type DetailedCart struct {
	ID                  int                    `json:"id"`
	CreatedAt           string                 `json:"created_at"`
	UpdatedAt           string                 `json:"updated_at"`
	ConvertedAt         string                 `json:"converted_at"`
	IsActive            bool                   `json:"is_active"`
	IsVirtual           bool                   `json:"is_virtual"`
	Items               []Item                 `json:"items"`
	ItemsCount          int                    `json:"items_count"`
	ItemsQty            int                    `json:"items_qty"`
	Customer            Customer               `json:"customer"`
	BillingAddress      Address                `json:"billing_address"`
	ReservedOrderID     int                    `json:"reserved_order_id"`
	OrigOrderID         int                    `json:"orig_order_id"`
	Currency            Currency               `json:"currency"`
	CustomerIsGuest     bool                   `json:"customer_is_guest"`
	CustomerNote        string                 `json:"customer_note"`
	CustomerNoteNotify  bool                   `json:"customer_note_notify"`
	CustomerTaxClassID  int                    `json:"customer_tax_class_id"`
	StoreID             int                    `json:"store_id"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes"`
}

type DownloadableOption struct {
	DownloadableLinks []int `json:"downloadable_links,omitempty"`
}

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

type FileInfo struct {
	Base64EncodedData string `json:"base64_encoded_data,omitempty"`
	Type              string `json:"type,omitempty"`
	Name              string `json:"name,omitempty"`
}

type GiftcardItemOption struct {
	GiftcardAmount         string                 `json:"giftcard_amount,omitempty"`
	CustomGiftcardAmount   int                    `json:"custom_giftcard_amount,omitempty"`
	GiftcardSenderName     string                 `json:"giftcard_sender_name,omitempty"`
	GiftcardRecipientName  string                 `json:"giftcard_recipient_name,omitempty"`
	GiftcardSenderEmail    string                 `json:"giftcard_sender_email,omitempty"`
	GiftcardRecipientEmail string                 `json:"giftcard_recipient_email,omitempty"`
	GiftcardMessage        string                 `json:"giftcard_message,omitempty"`
	ExtensionAttributes    map[string]interface{} `json:"extension_attributes,omitempty"`
}

type Item struct {
	ItemID              int                    `json:"item_id,omitempty"`
	Sku                 string                 `json:"sku"`
	Qty                 float64                `json:"qty"`
	Name                string                 `json:"name,omitempty"`
	Price               float64                `json:"price,omitempty"`
	ProductType         string                 `json:"product_type,omitempty"`
	QuoteID             string                 `json:"quote_id"`
	ProductOption       ProductOption          `json:"product_option,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type NegotiableQuoteItem struct {
	ItemID                 int                    `json:"item_id,omitempty"`
	OriginalPrice          int                    `json:"original_price,omitempty"`
	OriginalTaxAmount      int                    `json:"original_tax_amount,omitempty"`
	OriginalDiscountAmount int                    `json:"original_discount_amount,omitempty"`
	ExtensionAttributes    map[string]interface{} `json:"extension_attributes,omitempty"`
}

type PaymentMethod struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

type PaymentMethodCode struct {
	Method string `json:"method"`
}

type ProductOption struct {
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type Region struct {
	RegionCode          string                 `json:"region_code,omitempty"`
	Region              string                 `json:"region,omitempty"`
	RegionID            int                    `json:"region_id,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type Shipping struct {
	Address             Address                `json:"address,omitempty"`
	Method              string                 `json:"method,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type ShippingAssignments struct {
	Shipping            Shipping               `json:"shipping,omitempty"`
	Items               []Item                 `json:"items,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type Carrier struct {
	CarrierCode  string  `json:"carrier_code"`
	MethodCode   string  `json:"method_code"`
	CarrierTitle string  `json:"carrier_title"`
	MethodTitle  string  `json:"method_title"`
	Amount       float64 `json:"amount"`
	BaseAmount   float64 `json:"base_amount"`
	Available    bool    `json:"available"`
	ErrorMessage string  `json:"error_message,omitempty"`
	PriceExclTax float64 `json:"price_excl_tax"`
	PriceInclTax float64 `json:"price_incl_tax"`
}

type StatusHistory struct {
	Comment             string                 `json:"comment,omitempty"`
	CreatedAt           string                 `json:"created_at,omitempty"`
	EntityID            int                    `json:"entity_id,omitempty"`
	EntityName          string                 `json:"entity_name,omitempty"`
	IsCustomerNotified  int                    `json:"is_customer_notified,omitempty"`
	IsVisibleOnFront    int                    `json:"is_visible_on_front,omitempty"`
	ParentID            int                    `json:"parent_id,omitempty"`
	Status              string                 `json:"status,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
