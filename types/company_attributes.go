package types

type CompanyAttributes struct {
	CustomerID          int                    `json:"customer_id,omitempty"`
	CompanyID           int                    `json:"company_id,omitempty"`
	JobTitle            string                 `json:"job_title,omitempty"`
	Status              int                    `json:"status,omitempty"`
	Telephone           string                 `json:"telephone,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
