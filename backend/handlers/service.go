package handlers

import (
	"github.com/gin-gonic/gin"
	"go-backend/database"
	"go-backend/models"
	"net/http"
	"log"
)

// GetServices - ดึงข้อมูลบริการทั้งหมด
func GetServices(c *gin.Context) {
	rows, err := database.DB.Query("SELECT serviceID, name, type, description, price FROM \"Service\"")
	if err != nil {
		log.Println("Error fetching services:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch services"})
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		if err := rows.Scan(&service.ServiceID, &service.Name, &service.Type, &service.Description, &service.Price); err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process services"})
			return
		}
		services = append(services, service)
	}

	c.JSON(http.StatusOK, services)
}

// GetServiceByID - ดึงข้อมูลบริการตาม serviceID
func GetServiceByID(c *gin.Context) {
	serviceID := c.Param("serviceID")
	var service models.Service
	err := database.DB.QueryRow(`SELECT serviceID, name, type, description, price FROM "Service" WHERE serviceID = $1`, serviceID).
		Scan(&service.ServiceID, &service.Name, &service.Type, &service.Description, &service.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// CreateService - สร้างบริการใหม่
func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.QueryRow(
		`INSERT INTO "Service" (name, type, description, price) VALUES ($1, $2, $3, $4) RETURNING serviceID`,
		service.Name, service.Type, service.Description, service.Price).Scan(&service.ServiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}

	c.JSON(http.StatusCreated, service)
}

// EditService - แก้ไขบริการตาม serviceID
func EditService(c *gin.Context) {
	serviceID := c.Param("serviceID")
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูลบริการ
	_, err := database.DB.Exec(
		`UPDATE "Service" SET name = $1, type = $2, description = $3, price = $4 WHERE serviceID = $5`,
		service.Name, service.Type, service.Description, service.Price, serviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}

	// ดึงข้อมูลบริการที่อัปเดต
	err = database.DB.QueryRow(`SELECT serviceID, name, type, description, price FROM "Service" WHERE serviceID = $1`, serviceID).
		Scan(&service.ServiceID, &service.Name, &service.Type, &service.Description, &service.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// DeleteService - ลบบริการตาม serviceID
func DeleteService(c *gin.Context) {
	serviceID := c.Param("serviceID")

	// ลบบริการจากฐานข้อมูล
	_, err := database.DB.Exec(`DELETE FROM "Service" WHERE serviceID = $1`, serviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}
