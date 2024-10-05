package request

type PaymentRequest struct {
	UserID string `json:"user_id"`
	Amount uint64 `json:"amount"`
}
