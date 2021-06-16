package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func CreateDB() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }

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
	
	// DB.AutoMigrate(&models.User{}, &models.Trip{})
}
