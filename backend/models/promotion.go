package models

type Promotion struct {
    PromotionID uint    `json:"promotionID"`
    Code        string  `json:"code"`
    Discount    float64 `json:"discount"`
    Description string  `json:"description"`
}