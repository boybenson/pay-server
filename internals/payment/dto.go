package payment

type CreatePaymentCodeRequest struct {
	User string `json:"user" validate:"required"`
	Code string `json:"code" validate:"required"`
}
