package handlers

import (
	"github.com/gin-gonic/gin"
	"go-backend/database"
	"go-backend/models"
	"net/http"
)

// GetUsers - ดึงข้อมูลผู้ใช้ทั้งหมด
func GetUsers(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT userID, email, role FROM "users"`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID - ดึงข้อมูลผู้ใช้ตาม userID
func GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User
	err := database.DB.QueryRow(`SELECT userID, email, role FROM "users" WHERE userID = $1`, userID).
		Scan(&user.UserID, &user.Email, &user.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser - สร้างผู้ใช้ใหม่
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เช็คถ้าผู้ใช้มีอยู่แล้วโดยใช้ email แทน username
	var exists bool
	err := database.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM "users" WHERE email = $1)`, user.Email).Scan(&exists)
	if err != nil || exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// สร้างผู้ใช้ใหม่
	err = database.DB.QueryRow(
		`INSERT INTO "users" (email, role, password) VALUES ($1, $2, $3) RETURNING userID`,
		user.Email, user.Role, user.Password).Scan(&user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// EditUser - แก้ไขข้อมูลผู้ใช้ตาม userID
func EditUser(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูลผู้ใช้
	_, err := database.DB.Exec(
		`UPDATE "users" SET email = $1, role = $2, password = $3 WHERE userID = $4`,
		user.Email, user.Role, user.Password, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// ดึงข้อมูลผู้ใช้ที่อัปเดต
	err = database.DB.QueryRow(`SELECT userID, email, role FROM "users" WHERE userID = $1`, userID).
		Scan(&user.UserID, &user.Email, &user.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser - ลบผู้ใช้ตาม userID
func DeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	// ลบผู้ใช้จากฐานข้อมูล
	_, err := database.DB.Exec(`DELETE FROM "users" WHERE userID = $1`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
