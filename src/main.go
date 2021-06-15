package main

import (
	"fmt"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	
	database.CreateDB()
	fmt.Println(database.DB)
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yeeeee-haw!")
	})

	userGroup := e.Group("/users")
	tripGroup := e.Group("/trips")

	userGroup.GET("/", handlers.FindAllUsers)
	userGroup.POST("/", handlers.CreateUser)


	tripGroup.GET("/", handlers.FindAllTrips)
	tripGroup.POST("/", handlers.CreateTrip)
	tripGroup.POST("/test", handlers.Test)


	e.Logger.Fatal(e.Start(":8000"))
}
