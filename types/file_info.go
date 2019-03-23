package types

type FileInfo struct {
	Base64EncodedData string `json:"base64_encoded_data,omitempty"`
	Type              string `json:"type,omitempty"`
	Name              string `json:"name,omitempty"`
}
