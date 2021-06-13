package main

import (
	"fmt"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/models"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var users []models.User

func main() {
	
	CreateDB()
	fmt.Println(DB)
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yeeeee-haw!")
	})

	e.GET("/test", func(c echo.Context) error {
		user := models.User {
			FirstName: "Connor",
			LastName: "Jacobs",
			Email: "cjacob22@gmail.com",
			Phone: "314-540-4529",
			Location: models.UserLocation{
				City: "Denver",
				State: "CO",
				Country: "US",
			},
		}
		fmt.Println(&user)

		result := DB.Create(&user)
		fmt.Println("result.Error", result.Error)
		return c.String(http.StatusOK, "Created New Test User")
	})


	e.Logger.Fatal(e.Start(":8000"))
}
