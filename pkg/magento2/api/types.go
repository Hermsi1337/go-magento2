package api

const (
	Customer AuthenticationType = iota
	Administrator
)

type AuthenticationType int

func (authenticationType AuthenticationType) Route() string {
	return [...]string{integrationCustomerTokenService, integrationAdminTokenService}[authenticationType]
}

type AuthenticationRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
