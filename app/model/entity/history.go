package entity

type History struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Action     string `json:"action"`
	Amount     uint64 `json:"amount,omitempty"`
	Timestamp  string `json:"timestamp"`
}
