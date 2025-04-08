package handlers

import (
	"github.com/gin-gonic/gin"
	"go-backend/database"
	"go-backend/models"
	"net/http"
)

// GetPromotionsAll - ดึงข้อมูลโปรโมชั่นทั้งหมด
func GetPromotionsAll(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT promotionID, code, discount, description FROM "Promotion"`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch promotions"})
		return
	}
	defer rows.Close()

	var promotions []models.Promotion
	for rows.Next() {
		var p models.Promotion
		if err := rows.Scan(&p.PromotionID, &p.Code, &p.Discount, &p.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		promotions = append(promotions, p)
	}

	c.JSON(http.StatusOK, promotions)
}

// GetPromotionsByID - ดึงข้อมูลโปรโมชั่นตาม promotionID
func GetPromotionsByID(c *gin.Context) {
	promotionID := c.Param("promotionID")
	var promotion models.Promotion
	err := database.DB.QueryRow(`SELECT promotionID, code, discount, description FROM "Promotion" WHERE promotionID = $1`, promotionID).
		Scan(&promotion.PromotionID, &promotion.Code, &promotion.Discount, &promotion.Description)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
		return
	}

	c.JSON(http.StatusOK, promotion)
}

// CreatePromotion - สร้างโปรโมชั่นใหม่
func CreatePromotion(c *gin.Context) {
	var promotion models.Promotion
	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.QueryRow(
		`INSERT INTO "Promotion" (code, discount, description) VALUES ($1, $2, $3) RETURNING promotionID`,
		promotion.Code, promotion.Discount, promotion.Description).Scan(&promotion.PromotionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create promotion"})
		return
	}

	c.JSON(http.StatusCreated, promotion)
}

// EditPromotion - แก้ไขโปรโมชั่นตาม promotionID
func EditPromotion(c *gin.Context) {
	promotionID := c.Param("promotionID")
	var promotion models.Promotion
	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูลโปรโมชั่น
	_, err := database.DB.Exec(
		`UPDATE "Promotion" SET code = $1, discount = $2, description = $3 WHERE promotionID = $4`,
		promotion.Code, promotion.Discount, promotion.Description, promotionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update promotion"})
		return
	}

	// ดึงข้อมูลโปรโมชั่นที่อัปเดต
	err = database.DB.QueryRow(`SELECT promotionID, code, discount, description FROM "Promotion" WHERE promotionID = $1`, promotionID).
		Scan(&promotion.PromotionID, &promotion.Code, &promotion.Discount, &promotion.Description)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
		return
	}

	c.JSON(http.StatusOK, promotion)
}

// DeletePromotion - ลบโปรโมชั่นตาม promotionID
func DeletePromotion(c *gin.Context) {
	promotionID := c.Param("promotionID")

	// ลบโปรโมชั่นจากฐานข้อมูล
	_, err := database.DB.Exec(`DELETE FROM "Promotion" WHERE promotionID = $1`, promotionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete promotion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Promotion deleted successfully"})
}
