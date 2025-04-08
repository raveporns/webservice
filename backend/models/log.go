package models

type Log struct {
    LogID     uint            `json:"logID"`
    UserID    *uint           `json:"userID"` // nullable
    Action    string          `json:"action"`
    Target    string          `json:"target"`
    Timestamp string          `json:"timestamp"`
    Referrer  string          `json:"referrer"`
    Parameter map[string]any  `json:"parameter"` // JSON object
}
