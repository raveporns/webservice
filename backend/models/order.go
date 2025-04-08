package models

import "time"

type Order struct {
    OrderID     uint      `gorm:"primaryKey" json:"order_id"`
    UserID      uint      `json:"user_id"`
    ServiceID   uint      `json:"service_id"`
    PromotionID *uint     `json:"promotion_id"`
    CreatedAt   time.Time `json:"created_at"`
}
