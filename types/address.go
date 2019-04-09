package types

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
