package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	loadEnv()

	dsn := buildDSN()
	// dsn := "user=asistencia password=1234456 dbname=asistencia port=5432 sslmode=disable"
	// PGPASSWORD=7xt3Vx6eAevhZTMmSiGJ psql -h containers-us-west-210.railway.app -U postgres -p 7112 -d railway
	// dsn := "host=localhost user=asistencias password=123456 dbname=asistencias port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = db

	log.Println("Conexión establecida con Postgress")
}

// func newDB() (*gorm.DB, error){
// 	dsn := "user=postgres password=1234 dbname=smartdb port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}
// 	return db
// }

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func buildDSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	return fmt.Sprintf(dsn, dbHost, dbUser, dbPassword, dbName, dbPort)
}
