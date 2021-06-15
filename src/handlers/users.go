package handlers

import (
	"fmt"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := models.User{
		FirstName: "Connor",
		LastName:  "Jacobs",
		Email:     "cjacob22@gmail.com",
		Phone:     "314-540-4529",
		Location:  models.UserLocation{City: "Denver", State: "CO", Country: "US"},
		Trips:     []*models.Trip{},
	}
	fmt.Println(&user)

	result := database.DB.Create(&user)
	fmt.Println("result.Error", result.Error)
	return c.String(http.StatusOK, "Created New Test User")
}

func FindAllUsers(c echo.Context) error {
	var user models.User
	database.DB.Find(&user)
	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}