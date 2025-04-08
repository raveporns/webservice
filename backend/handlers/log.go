package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-backend/database"
	"go-backend/models"
	"net/http"
)

// CreateLog - สร้าง log ใหม่
func CreateLog(c *gin.Context) {
	var logEntry models.Log
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แปลง parameter เป็น JSON
	parameterJSON, err := json.Marshal(logEntry.Parameter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format"})
		return
	}

	_, err = database.DB.Exec(`INSERT INTO "Log" (userID, action, target, referrer, parameter) 
        VALUES ($1, $2, $3, $4, $5)`,
		logEntry.UserID, logEntry.Action, logEntry.Target, logEntry.Referrer, parameterJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Log recorded"})
}

// GetLogs - ดึงข้อมูลทั้งหมดของ logs
func GetLogs(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT logID, userID, action, target, referrer, parameter FROM "Log"`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch logs"})
		return
	}
	defer rows.Close()

	var logs []models.Log
	for rows.Next() {
		var logEntry models.Log
		var parameterJSON []byte
		if err := rows.Scan(&logEntry.LogID, &logEntry.UserID, &logEntry.Action, &logEntry.Target, &logEntry.Referrer, &parameterJSON); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// แปลง parameter จาก JSON กลับเป็นแผนที่
		if err := json.Unmarshal(parameterJSON, &logEntry.Parameter); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal parameter"})
			return
		}

		logs = append(logs, logEntry)
	}

	c.JSON(http.StatusOK, logs)
}

// GetLogByID - ดึงข้อมูล log ตาม logID
func GetLogByID(c *gin.Context) {
	logID := c.Param("logID")
	var logEntry models.Log
	var parameterJSON []byte

	err := database.DB.QueryRow(`SELECT logID, userID, action, target, referrer, parameter FROM "Log" WHERE logID = $1`, logID).
		Scan(&logEntry.LogID, &logEntry.UserID, &logEntry.Action, &logEntry.Target, &logEntry.Referrer, &parameterJSON)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	// แปลง parameter จาก JSON กลับเป็นแผนที่
	if err := json.Unmarshal(parameterJSON, &logEntry.Parameter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal parameter"})
		return
	}

	c.JSON(http.StatusOK, logEntry)
}

// EditLog - แก้ไขข้อมูล log ตาม logID
func EditLog(c *gin.Context) {
	logID := c.Param("logID")
	var logEntry models.Log
	if err := c.ShouldBindJSON(&logEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แปลง parameter เป็น JSON
	parameterJSON, err := json.Marshal(logEntry.Parameter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format"})
		return
	}

	// อัปเดตข้อมูลในฐานข้อมูล
	_, err = database.DB.Exec(`UPDATE "Log" SET userID = $1, action = $2, target = $3, referrer = $4, parameter = $5 
		WHERE logID = $6`,
		logEntry.UserID, logEntry.Action, logEntry.Target, logEntry.Referrer, parameterJSON, logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log updated"})
}

// DeleteLog - ลบ log ตาม logID
func DeleteLog(c *gin.Context) {
	logID := c.Param("logID")

	// ลบข้อมูล log ที่มี logID
	_, err := database.DB.Exec(`DELETE FROM "Log" WHERE logID = $1`, logID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log deleted successfully"})
}
