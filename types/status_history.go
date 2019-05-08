package types

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
