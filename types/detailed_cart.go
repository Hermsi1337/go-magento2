package types

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
