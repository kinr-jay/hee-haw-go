package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(&user)
	
	hashbytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashbytes)

	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, "Unable to create new account.")
	}
	fmt.Println("user ID = ", user.ID)
	token, err2 := CreateJWTToken(user.Email, user.ID)
	if err2 != nil {
		fmt.Println("Error while creating JWT token: ", err2)
		return c.JSON(http.StatusInternalServerError, "Unable to login.")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful.",
		"token": token,
	})
}

func FindAllUsers(c echo.Context) error {
	var user models.User
	database.DB.Find(&user)
	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}

