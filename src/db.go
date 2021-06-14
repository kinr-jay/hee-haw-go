package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kinr-jay/hee-haw-go/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func CreateDB() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }


	// HOST := os.Getenv("HOST")
	// PGUSER := os.Getenv("PGUSER")
	// PGPASSWORD := os.Getenv("PGPASSWORD")
	// DBNAME := os.Getenv("DBNAME")
	// PORT := os.Getenv("PORT")
	// SSLMODE := os.Getenv("SSLMODE")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, PGUSER, PGPASSWORD, DBNAME, PORT, SSLMODE)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
    log.Fatal(err)
  } else {
		fmt.Println("Connected to the DB")
	}
	DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
  if err != nil {
    log.Fatal(err)
  } else {
		fmt.Println("Converted DB connection to GORM")
	}

	DB.AutoMigrate(&models.User{}, &models.Trip{})
}
