package magento2

type AuthenticationRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}