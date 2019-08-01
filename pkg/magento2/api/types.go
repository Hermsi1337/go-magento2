package api

const (
	AnonymousClientType     = "anonymous"
	CustomerClientType      = "customer"
	AdministratorClientType = "administrator"
)

type AuthenticationRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
