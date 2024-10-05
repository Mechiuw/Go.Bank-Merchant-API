package entity

type Merchant struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	Balance     uint64    `json:"balance"`
	AccountId   Account   `json:"account_id"`
	HistoryList []History `json:"history_list"`
}
