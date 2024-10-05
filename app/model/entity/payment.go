package entity

type Payment struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	MerchantID string `json:"merchant_id"`
	Amount     uint64 `json:"amount"`
	Timestamp  string `json:"timestamp"`
}
