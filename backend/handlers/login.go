package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// เก็บข้อมูลผู้ใช้ (ตัวอย่างในหน่วยความจำ)
var users = map[string]string{} // key: username, value: hashed password

// Hash รหัสผ่าน
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// ตรวจสอบชื่อผู้ใช้และรหัสผ่าน (เก็บรหัสผ่านในรูปแบบแฮช)
	storedHashedPassword, exists := users[loginData.Username]
	if !exists || !checkPasswordHash(loginData.Password, storedHashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// สร้าง JWT ให้กับผู้ใช้
	token, err := generateToken(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// ส่ง Token กลับไปให้ Client
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ตรวจสอบรหัสผ่าน
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}