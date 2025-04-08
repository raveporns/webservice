package handlers

import (
	"github.com/gin-gonic/gin"
	"go-backend/database"
	"go-backend/models"
	"net/http"
)

// GetSchedules - ดึงข้อมูลทั้งหมดของ schedule
func GetSchedules(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT scheduleID, orderID, scheduleDate FROM "Schedule"`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var schedule models.Schedule
		if err := rows.Scan(&schedule.ScheduleID, &schedule.OrderID, &schedule.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		schedules = append(schedules, schedule)
	}

	c.JSON(http.StatusOK, schedules)
}

// GetScheduleByID - ดึงข้อมูลตาม scheduleID
func GetScheduleByID(c *gin.Context) {
	scheduleID := c.Param("scheduleID")
	var schedule models.Schedule
	err := database.DB.QueryRow(`SELECT scheduleID, orderID, scheduleDate FROM "Schedule" WHERE scheduleID = $1`, scheduleID).
		Scan(&schedule.ScheduleID, &schedule.OrderID, &schedule.Date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// CreateSchedule - สร้าง schedule ใหม่
func CreateSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id uint
	err := database.DB.QueryRow(`INSERT INTO "Schedule" (orderID, scheduleDate) 
		VALUES ($1, $2) RETURNING scheduleID`, schedule.OrderID, schedule.Date).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	schedule.ScheduleID = id
	c.JSON(http.StatusCreated, schedule)
}

// EditSchedule - แก้ไขข้อมูล schedule ตาม scheduleID
func EditSchedule(c *gin.Context) {
	scheduleID := c.Param("scheduleID")
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูล schedule
	_, err := database.DB.Exec(
		`UPDATE "Schedule" SET orderID = $1, scheduleDate = $2 WHERE scheduleID = $3`,
		schedule.OrderID, schedule.Date, scheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	// ดึงข้อมูล schedule ที่อัปเดต
	err = database.DB.QueryRow(`SELECT scheduleID, orderID, scheduleDate FROM "Schedule" WHERE scheduleID = $1`, scheduleID).
		Scan(&schedule.ScheduleID, &schedule.OrderID, &schedule.Date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// DeleteSchedule - ลบ schedule ตาม scheduleID
func DeleteSchedule(c *gin.Context) {
	scheduleID := c.Param("scheduleID")

	// ลบ schedule
	_, err := database.DB.Exec(`DELETE FROM "Schedule" WHERE scheduleID = $1`, scheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}
