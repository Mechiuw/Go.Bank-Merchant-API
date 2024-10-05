package entity

type Merchant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Balance uint64 `json:"balance"`
}
