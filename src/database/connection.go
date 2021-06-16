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
  //////// Load .env file if in Development environment ///////////
	if os.Getenv("PRODUCTION") == "false" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//////// Connect to Heroku's Postgres Database //////////
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
	
	//////// Auto-Migrate Tables to Database ///////////
	// DB.AutoMigrate(&models.User{}, &models.Trip{})
}
