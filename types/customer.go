package types

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
