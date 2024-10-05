package response

type HistoryResponse struct {
	Histories []HistoryEntry `json:"histories"` // Array of history entries
}

type HistoryEntry struct {
	ID        string  `json:"id"`               // Unique identifier for the log entry
	Action    string  `json:"action"`           // Description of the activity
	Amount    float64 `json:"amount,omitempty"` // Amount involved in the action, if applicable
	Timestamp string  `json:"timestamp"`        // Time of the action
}
