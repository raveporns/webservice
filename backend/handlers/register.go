package handlers

import (

	"go-backend/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)



func Register(c *gin.Context) {
    if database.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not established"})
        return
    }

    var user struct {
        Email    string `json:"email"`
        Password string `json:"password"`
		Role     string `json:"role"`
    }

    // รับข้อมูลจาก Request Body
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // เข้ารหัส password ก่อนเก็บลงฐานข้อมูล
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // เพิ่มผู้ใช้ใหม่ลงในฐานข้อมูล
    insertQuery := "INSERT INTO \"users\" (email, password, role) VALUES ($1, $2, $3)"
    _, err = database.DB.Exec(insertQuery, user.Email, hashedPassword, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
        return
    }

    // สร้าง token สำหรับผู้ใช้ใหม่
    token, err := generateToken(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}


// Secret key สำหรับใช้เข้ารหัส Token (ควรเก็บไว้เป็นความลับ)
var jwtSecret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAYWRtaW4uY29tIiwiZXhwIjoxNzQ0MzkzNDAxfQ.iqlX2B5_QeloLtx5pCga4jYi70x4Ztcsv8he0Mwlbi4")

// generateToken สร้าง JWT โดยรับ email เป็น claim พร้อมตั้งเวลาหมดอายุ (exp)
func generateToken(email string) (string, error) {
	// สร้าง token พร้อมกำหนด claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(72 * time.Hour).Unix(), // กำหนดอายุ token 72 ชั่วโมง
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ลงนาม token ด้วย secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
