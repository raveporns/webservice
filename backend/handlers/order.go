package handlers

import (
    "github.com/gin-gonic/gin"
    "go-backend/database"
    "go-backend/models"
    "net/http"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var orderID int64
    err := database.DB.QueryRow("INSERT INTO \"Order\" (userID, serviceID, promotionID) VALUES ($1, $2, $3) RETURNING orderID", order.UserID, order.ServiceID, order.PromotionID).Scan(&orderID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    order.OrderID = uint(orderID)
    c.JSON(http.StatusCreated, order)
}

func GetOrderAll(c *gin.Context) {
	var orders []models.Order
	rows, err := database.DB.Query("SELECT orderID, userID, serviceID, promotionID FROM \"Order\"")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.ServiceID, &order.PromotionID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}


func GetOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")
	var order models.Order
	err := database.DB.QueryRow("SELECT orderID, userID, serviceID, promotionID FROM \"Order\" WHERE orderID = $1", orderID).Scan(&order.OrderID, &order.UserID, &order.ServiceID, &order.PromotionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}


func EditOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แก้ไขข้อมูลในฐานข้อมูล
	_, err := database.DB.Exec("UPDATE \"Order\" SET userID = $1, serviceID = $2, promotionID = $3 WHERE orderID = $4", order.UserID, order.ServiceID, order.PromotionID, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ดึงข้อมูลออเดอร์ที่ถูกแก้ไข
	err = database.DB.QueryRow("SELECT orderID, userID, serviceID, promotionID FROM \"Order\" WHERE orderID = $1", orderID).Scan(&order.OrderID, &order.UserID, &order.ServiceID, &order.PromotionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}


func DeleteOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")

	// ลบข้อมูลในฐานข้อมูล
	_, err := database.DB.Exec("DELETE FROM \"Order\" WHERE orderID = $1", orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
