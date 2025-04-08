// models/service.go
package models

type Service struct {
    ServiceID   uint    `json:"serviceID" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Type        string  `json:"type"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}
