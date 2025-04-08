package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var DB *sql.DB

// Connect ฟังก์ชันเชื่อมต่อฐานข้อมูล PostgreSQL
func Connect() {
	// โหลดไฟล์ .env
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using default values")
	}

	// ตรวจสอบค่าจาก environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	// สร้าง DSN สำหรับเชื่อมต่อกับฐานข้อมูล
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	
	// เชื่อมต่อกับฐานข้อมูล PostgreSQL
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// ตรวจสอบการเชื่อมต่อ
	if err := DB.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// ปรับแต่งการตั้งค่าการเชื่อมต่อฐานข้อมูล
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(time.Hour)

	// แสดงข้อความว่าการเชื่อมต่อสำเร็จ
	log.Println("Successfully connected to PostgreSQL!")
}
