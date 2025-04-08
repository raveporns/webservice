package models

type Schedule struct {
    ScheduleID  uint   `json:"scheduleID"`
    Date        string `json:"scheduleDate"` // ISO format (e.g. 2025-04-10T14:00:00Z)
    OrderID     uint   `json:"orderID"`
}