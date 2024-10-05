package entity

type Transaction struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Amount    uint64 `json:"amount"`
	Timestamp string `json:"timestamp"`
}
