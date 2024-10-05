package response

type HistoryResponse struct {
	Histories []HistoryEntry `json:"histories"`
}

type HistoryEntry struct {
	ID        string  `json:"id"`
	Action    string  `json:"action"`
	Amount    float64 `json:"amount,omitempty"`
	Timestamp string  `json:"timestamp"`
}
