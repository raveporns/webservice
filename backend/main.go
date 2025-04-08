package main

import (
	"fmt"
	"go-backend/database"
	"go-backend/handlers"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func main() {

	database.Connect()
	
	r := gin.Default()

	// ใช้ CORS Middleware
	r.Use(cors.Default())

	// Endpoint สำหรับการลงทะเบียน (register)
	r.POST("/register", handlers.Register)
	r.GET("/user", handlers.GetUsers)

	// Endpoint สำหรับการเข้าสู่ระบบ (login)
	r.POST("/login", handlers.Login)

	r.GET("/service", handlers.GetServices)
	r.GET("/service/:id", handlers.GetServiceByID)
	r.POST("/service", handlers.CreateService)
	r.PUT("/service/:id", handlers.EditService)
	r.DELETE("/service/:id", handlers.DeleteService)

	r.GET("/promotion", handlers.GetUsers)
	r.GET("/promotion/:id", handlers.GetUserByID)
	r.POST("/promotion/", handlers.CreateUser)
	r.PUT("/promotion/:id", handlers.EditUser)
	r.DELETE("/promotion/:id", handlers.DeleteUser)

	// ใช้ JWTAuthMiddleware สำหรับ Route ที่ต้องการการยืนยันตัวตน
	protected := r.Group("/protected")
	protected.Use(JWTAuthMiddleware())
	{
		// ตัวอย่าง endpoint ที่ต้องการการยืนยันตัวตน
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Welcome %s! You have accessed a protected endpoint.", username),
			})
		})

		protected.GET("/order", handlers.GetOrderAll)
		protected.GET("/order/:id", handlers.GetOrderByID)
		protected.POST("/order", handlers.CreateOrder)
		protected.PUT("/order/:id", handlers.EditOrderByID)
		protected.DELETE("/order/:id", handlers.DeleteOrderByID)

		protected.GET("/user", handlers.GetUsers)
		protected.GET("/user/:id", handlers.GetUserByID)
		// protected.POST("/register", handlers.CreateUser)
		protected.PUT("/user/:id", handlers.EditUser)
		protected.DELETE("/user/:id", handlers.DeleteUser)

		protected.GET("/schedule", handlers.GetSchedules)
		protected.GET("/schedule/:id", handlers.GetScheduleByID)
		protected.POST("/schedule", handlers.CreateSchedule)
		protected.PUT("/schedule/:id", handlers.EditSchedule)
		protected.DELETE("/schedule/:id", handlers.DeleteSchedule)

	}

	// เริ่ม API ที่พอร์ต 8082
	r.Run(":8082")
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// JWTAuthMiddleware เป็น middleware สำหรับตรวจสอบ JWT ใน Header ของ Request
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		var tokenString string
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
		}
		c.Next()
	}
}


