package cart

import "fmt"

type ItemNotFoundError struct {
	ItemID int
}

func (a *ItemNotFoundError) Error() string {
	return fmt.Sprintf("itemID '%d' is non-existent", a.ItemID)
}

type Cart struct {
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

type ProductOption struct {
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}

type Customer struct {
	ID                     int                      `json:"id,omitempty"`
	GroupID                int                      `json:"group_id,omitempty"`
	DefaultBilling         string                   `json:"default_billing,omitempty"`
	DefaultShipping        string                   `json:"default_shipping,omitempty"`
	Confirmation           string                   `json:"confirmation,omitempty"`
	CreatedAt              string                   `json:"created_at,omitempty"`
	UpdatedAt              string                   `json:"updated_at,omitempty"`
	CreatedIn              string                   `json:"created_in,omitempty"`
	Dob                    string                   `json:"dob,omitempty"`
	Email                  string                   `json:"email,omitempty"`
	Firstname              string                   `json:"firstname,omitempty"`
	Lastname               string                   `json:"lastname,omitempty"`
	Middlename             string                   `json:"middlename,omitempty"`
	Prefix                 string                   `json:"prefix,omitempty"`
	Suffix                 string                   `json:"suffix,omitempty"`
	Gender                 int                      `json:"gender,omitempty"`
	StoreID                int                      `json:"store_id,omitempty"`
	Taxvat                 string                   `json:"taxvat,omitempty"`
	WebsiteID              int                      `json:"website_id,omitempty"`
	Addresses              []Address                `json:"addresses,omitempty"`
	DisableAutoGroupChange int                      `json:"disable_auto_group_change,omitempty"`
	ExtensionAttributes    map[string]interface{}   `json:"extension_attributes,omitempty"`
	CustomAttributes       []map[string]interface{} `json:"custom_attributes,omitempty"`
}

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

type AddressInformation struct {
	ShippingAddress      Address                  `json:"shipping_address"`
	BillingAddress       Address                  `json:"billing_address"`
	ShippingMethodCode   string                   `json:"shipping_method_code"`
	ShippingCarrierCodes string                   `json:"shipping_carrier_code"`
	ExtensionAttributes  map[string]interface{}   `json:"extension_attributes,omitempty"`
	CustomAttributes     []map[string]interface{} `json:"custom_attributes,omitempty"`
}

type PaymentMethodCode struct {
	Method string `json:"method"`
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

type PaymentMethod struct {
	Code  string `json:"code"`
	Title string `json:"title"`
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
