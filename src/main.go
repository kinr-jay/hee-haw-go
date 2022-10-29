package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//////// Load .env file if in Development environment ///////////
	if os.Getenv("PRODUCTION") == "false" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	
	//////// Create connection to database, instantiate echo //////////
	database.CreateDB()
	e := echo.New()
	e.Use(middleware.CORS())
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yeeeee-haw!")
	})

	////// base url paths //////////
	e.POST("/login", handlers.Login)
	// e.POST("/check-auth", handlers.CheckJWT)
	e.POST("/signup", handlers.CreateUser)

	////// Echo Routing Groups /////////
	userGroup := e.Group("/users")
	tripGroup := e.Group("/trips")

	/////// JWT Middleware //////////////
	JWT_KEY := os.Getenv("JWT_KEY")
	userGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(JWT_KEY),
	}))
	tripGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(JWT_KEY),
	}))

	////// Users Routing /////////
	userGroup.GET("", handlers.FindAllUsers)
	userGroup.GET("/:userId", handlers.FindUser)
	userGroup.PUT("/:userId", handlers.UpdateUser)
	userGroup.DELETE("/:userId", handlers.DeleteUser)

	////// Trips Routing /////////
	tripGroup.GET("", handlers.FindAllTrips)
	tripGroup.POST("", handlers.CreateTrip)
	tripGroup.GET("/:tripId", handlers.FindTrip)
	tripGroup.PUT("/:tripId", handlers.UpdateTrip)
	tripGroup.DELETE("/:tripId", handlers.DeleteTrip)
	tripGroup.PUT("/:tripId/add/:userId", handlers.AddTeamMember)
	tripGroup.PUT("/:tripId/remove/:userId", handlers.RemoveTeamMember)

	/////// Start Echo Server //////////
	PORT := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(PORT))
}
