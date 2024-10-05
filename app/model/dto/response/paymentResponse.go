package response

type PaymentResponse struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}
