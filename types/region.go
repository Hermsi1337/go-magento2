package types

type Region struct {
	RegionCode          string                 `json:"region_code,omitempty"`
	Region              string                 `json:"region,omitempty"`
	RegionID            int                    `json:"region_id,omitempty"`
	ExtensionAttributes map[string]interface{} `json:"extension_attributes,omitempty"`
}
