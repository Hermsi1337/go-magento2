package types

type StatusHistory struct {
	Comment             string              `json:"comment"`
	CreatedAt           string              `json:"created_at"`
	EntityID            int                 `json:"entity_id"`
	EntityName          string              `json:"entity_name"`
	IsCustomerNotified  int                 `json:"is_customer_notified"`
	IsVisibleOnFront    int                 `json:"is_visible_on_front"`
	ParentID            int                 `json:"parent_id"`
	Status              string              `json:"status"`
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}
