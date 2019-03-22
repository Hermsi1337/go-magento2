package magento2

type AuthenticationRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DetailedCart struct {
	ID                  int                 `json:"id"`
	CreatedAt           string              `json:"created_at"`
	UpdatedAt           string              `json:"updated_at"`
	ConvertedAt         string              `json:"converted_at"`
	IsActive            bool                `json:"is_active"`
	IsVirtual           bool                `json:"is_virtual"`
	Items               []Items             `json:"items"`
	ItemsCount          int                 `json:"items_count"`
	ItemsQty            int                 `json:"items_qty"`
	Customer            Customer            `json:"customer"`
	BillingAddress      BillingAddress      `json:"billing_address"`
	ReservedOrderID     int                 `json:"reserved_order_id"`
	OrigOrderID         int                 `json:"orig_order_id"`
	Currency            Currency            `json:"currency"`
	CustomerIsGuest     bool                `json:"customer_is_guest"`
	CustomerNote        string              `json:"customer_note"`
	CustomerNoteNotify  bool                `json:"customer_note_notify"`
	CustomerTaxClassID  int                 `json:"customer_tax_class_id"`
	StoreID             int                 `json:"store_id"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type FileInfo struct {
	Base64EncodedData string `json:"base64_encoded_data"`
	Type              string `json:"type"`
	Name              string `json:"name"`
}

type ExtensionAttributes struct {
	FileInfo                FileInfo                  `json:"file_info,omitempty"`
	ShippingAssignments     []ShippingAssignments     `json:"shipping_assignments,omitempty"`
	NegotiableQuote         NegotiableQuote           `json:"negotiable_quote,omitempty"`
	CustomOptions           []CustomOptions           `json:"custom_options,omitempty"`
	BundleOptions           []BundleOptions           `json:"bundle_options,omitempty"`
	DownloadableOption      DownloadableOption        `json:"downloadable_option,omitempty"`
	GiftcardItemOption      GiftcardItemOption        `json:"giftcard_item_option,omitempty"`
	ConfigurableItemOptions []ConfigurableItemOptions `json:"configurable_item_options,omitempty"`
	NegotiableQuoteItem     NegotiableQuoteItem       `json:"negotiable_quote_item,omitempty"`
	CompanyAttributes       CompanyAttributes         `json:"company_attributes,omitempty"`
	IsSubscribed            bool                      `json:"is_subscribed,omitempty"`
	GiftRegistryID          int                       `json:"gift_registry_id,omitempty"`
}

type CustomOptions struct {
	OptionID            string              `json:"option_id"`
	OptionValue         string              `json:"option_value"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type BundleOptions struct {
	OptionID            int                 `json:"option_id"`
	OptionQty           int                 `json:"option_qty"`
	OptionSelections    []int               `json:"option_selections"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type DownloadableOption struct {
	DownloadableLinks []int `json:"downloadable_links"`
}

type GiftcardItemOption struct {
	GiftcardAmount         string              `json:"giftcard_amount"`
	CustomGiftcardAmount   int                 `json:"custom_giftcard_amount"`
	GiftcardSenderName     string              `json:"giftcard_sender_name"`
	GiftcardRecipientName  string              `json:"giftcard_recipient_name"`
	GiftcardSenderEmail    string              `json:"giftcard_sender_email"`
	GiftcardRecipientEmail string              `json:"giftcard_recipient_email"`
	GiftcardMessage        string              `json:"giftcard_message"`
	ExtensionAttributes    ExtensionAttributes `json:"extension_attributes"`
}

type ConfigurableItemOptions struct {
	OptionID            string              `json:"option_id"`
	OptionValue         int                 `json:"option_value"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type ProductOption struct {
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type NegotiableQuoteItem struct {
	ItemID                 int                 `json:"item_id"`
	OriginalPrice          int                 `json:"original_price"`
	OriginalTaxAmount      int                 `json:"original_tax_amount"`
	OriginalDiscountAmount int                 `json:"original_discount_amount"`
	ExtensionAttributes    ExtensionAttributes `json:"extension_attributes"`
}

type Items struct {
	ItemID              int                 `json:"item_id"`
	Sku                 string              `json:"sku"`
	Qty                 int                 `json:"qty"`
	Name                string              `json:"name"`
	Price               int                 `json:"price"`
	ProductType         string              `json:"product_type"`
	QuoteID             string              `json:"quote_id"`
	ProductOption       ProductOption       `json:"product_option"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type Region struct {
	RegionCode          string              `json:"region_code"`
	Region              string              `json:"region"`
	RegionID            int                 `json:"region_id"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type CustomAttributes struct {
	AttributeCode string `json:"attribute_code"`
	Value         string `json:"value"`
}

type Addresses struct {
	ID                  int                 `json:"id"`
	CustomerID          int                 `json:"customer_id"`
	Region              Region              `json:"region"`
	RegionID            int                 `json:"region_id"`
	CountryID           string              `json:"country_id"`
	Street              []string            `json:"street"`
	Company             string              `json:"company"`
	Telephone           string              `json:"telephone"`
	Fax                 string              `json:"fax"`
	Postcode            string              `json:"postcode"`
	City                string              `json:"city"`
	Firstname           string              `json:"firstname"`
	Lastname            string              `json:"lastname"`
	Middlename          string              `json:"middlename"`
	Prefix              string              `json:"prefix"`
	Suffix              string              `json:"suffix"`
	VatID               string              `json:"vat_id"`
	DefaultShipping     bool                `json:"default_shipping"`
	DefaultBilling      bool                `json:"default_billing"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
	CustomAttributes    []CustomAttributes  `json:"custom_attributes"`
}

type CompanyAttributes struct {
	CustomerID          int                 `json:"customer_id"`
	CompanyID           int                 `json:"company_id"`
	JobTitle            string              `json:"job_title"`
	Status              int                 `json:"status"`
	Telephone           string              `json:"telephone"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type Customer struct {
	ID                     int                 `json:"id"`
	GroupID                int                 `json:"group_id"`
	DefaultBilling         string              `json:"default_billing"`
	DefaultShipping        string              `json:"default_shipping"`
	Confirmation           string              `json:"confirmation"`
	CreatedAt              string              `json:"created_at"`
	UpdatedAt              string              `json:"updated_at"`
	CreatedIn              string              `json:"created_in"`
	Dob                    string              `json:"dob"`
	Email                  string              `json:"email"`
	Firstname              string              `json:"firstname"`
	Lastname               string              `json:"lastname"`
	Middlename             string              `json:"middlename"`
	Prefix                 string              `json:"prefix"`
	Suffix                 string              `json:"suffix"`
	Gender                 int                 `json:"gender"`
	StoreID                int                 `json:"store_id"`
	Taxvat                 string              `json:"taxvat"`
	WebsiteID              int                 `json:"website_id"`
	Addresses              []Addresses         `json:"addresses"`
	DisableAutoGroupChange int                 `json:"disable_auto_group_change"`
	ExtensionAttributes    ExtensionAttributes `json:"extension_attributes"`
	CustomAttributes       []CustomAttributes  `json:"custom_attributes"`
}

type BillingAddress struct {
	ID                  int                 `json:"id"`
	Region              string              `json:"region"`
	RegionID            int                 `json:"region_id"`
	RegionCode          string              `json:"region_code"`
	CountryID           string              `json:"country_id"`
	Street              []string            `json:"street"`
	Company             string              `json:"company"`
	Telephone           string              `json:"telephone"`
	Fax                 string              `json:"fax"`
	Postcode            string              `json:"postcode"`
	City                string              `json:"city"`
	Firstname           string              `json:"firstname"`
	Lastname            string              `json:"lastname"`
	Middlename          string              `json:"middlename"`
	Prefix              string              `json:"prefix"`
	Suffix              string              `json:"suffix"`
	VatID               string              `json:"vat_id"`
	CustomerID          int                 `json:"customer_id"`
	Email               string              `json:"email"`
	SameAsBilling       int                 `json:"same_as_billing"`
	CustomerAddressID   int                 `json:"customer_address_id"`
	SaveInAddressBook   int                 `json:"save_in_address_book"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
	CustomAttributes    []CustomAttributes  `json:"custom_attributes"`
}

type Currency struct {
	GlobalCurrencyCode  string              `json:"global_currency_code"`
	BaseCurrencyCode    string              `json:"base_currency_code"`
	StoreCurrencyCode   string              `json:"store_currency_code"`
	QuoteCurrencyCode   string              `json:"quote_currency_code"`
	StoreToBaseRate     int                 `json:"store_to_base_rate"`
	StoreToQuoteRate    int                 `json:"store_to_quote_rate"`
	BaseToGlobalRate    int                 `json:"base_to_global_rate"`
	BaseToQuoteRate     int                 `json:"base_to_quote_rate"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type Address struct {
	ID                  int                 `json:"id"`
	Region              string              `json:"region"`
	RegionID            int                 `json:"region_id"`
	RegionCode          string              `json:"region_code"`
	CountryID           string              `json:"country_id"`
	Street              []string            `json:"street"`
	Company             string              `json:"company"`
	Telephone           string              `json:"telephone"`
	Fax                 string              `json:"fax"`
	Postcode            string              `json:"postcode"`
	City                string              `json:"city"`
	Firstname           string              `json:"firstname"`
	Lastname            string              `json:"lastname"`
	Middlename          string              `json:"middlename"`
	Prefix              string              `json:"prefix"`
	Suffix              string              `json:"suffix"`
	VatID               string              `json:"vat_id"`
	CustomerID          int                 `json:"customer_id"`
	Email               string              `json:"email"`
	SameAsBilling       int                 `json:"same_as_billing"`
	CustomerAddressID   int                 `json:"customer_address_id"`
	SaveInAddressBook   int                 `json:"save_in_address_book"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
	CustomAttributes    []CustomAttributes  `json:"custom_attributes"`
}

type Shipping struct {
	Address             Address             `json:"address"`
	Method              string              `json:"method"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type ShippingAssignments struct {
	Shipping            Shipping            `json:"shipping"`
	Items               []Items             `json:"items"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type NegotiableQuote struct {
	QuoteID                  int                 `json:"quote_id"`
	IsRegularQuote           bool                `json:"is_regular_quote"`
	Status                   string              `json:"status"`
	NegotiatedPriceType      int                 `json:"negotiated_price_type"`
	NegotiatedPriceValue     int                 `json:"negotiated_price_value"`
	ShippingPrice            int                 `json:"shipping_price"`
	QuoteName                string              `json:"quote_name"`
	ExpirationPeriod         string              `json:"expiration_period"`
	EmailNotificationStatus  int                 `json:"email_notification_status"`
	HasUnconfirmedChanges    bool                `json:"has_unconfirmed_changes"`
	IsShippingTaxChanged     bool                `json:"is_shipping_tax_changed"`
	IsCustomerPriceChanged   bool                `json:"is_customer_price_changed"`
	Notifications            int                 `json:"notifications"`
	AppliedRuleIds           string              `json:"applied_rule_ids"`
	IsAddressDraft           bool                `json:"is_address_draft"`
	DeletedSku               string              `json:"deleted_sku"`
	CreatorID                int                 `json:"creator_id"`
	CreatorType              int                 `json:"creator_type"`
	OriginalTotalPrice       int                 `json:"original_total_price"`
	BaseOriginalTotalPrice   int                 `json:"base_original_total_price"`
	NegotiatedTotalPrice     int                 `json:"negotiated_total_price"`
	BaseNegotiatedTotalPrice int                 `json:"base_negotiated_total_price"`
	ExtensionAttributes      ExtensionAttributes `json:"extension_attributes"`
}
