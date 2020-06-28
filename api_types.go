package magento2

const (
	CustomerAuth AuthenticationType = iota
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
