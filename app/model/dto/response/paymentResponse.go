package response

type PaymentResponse struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}
