package types

type PaymentMethod struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

type PaymentMethodCode struct {
	Method string `json:"method"`
}
