package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"
	"gorm.io/gorm"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Create a new User Account
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
	token, err2 := CreateJWT(user.Email, user.ID)
	if err2 != nil {
		fmt.Println("Error while creating JWT token: ", err2)
		return c.JSON(http.StatusInternalServerError, "Unable to login.")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful.",
		"userId": user.ID,
		"token": token,
	})
}

// FindAllUsers for development only
func FindAllUsers(c echo.Context) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// Find Individual User Account by ID
func FindUser(c echo.Context) error {
	user := new(models.User)
	userId := c.Param("userId")
	result := database.DB.Preload("Trips", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "title")
	}).Select("id", "first_name", "last_name", "email", "phone").First(&user, userId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, "Could not find user account.")
	}
	return c.JSON(http.StatusOK, user)
}

// Update User Account
func UpdateUser(c echo.Context) error {
	user := new(models.User)
	updateId := c.Param("userId")
	result := database.DB.First(&user, updateId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, "Could not find user account.")
	}
	c.Bind(&user)
	result2 := database.DB.Save(&user)
	if result2.Error != nil {
		log.Fatal(result2.Error)
		return c.JSON(http.StatusInternalServerError, "Account update error.")
	}
	return c.JSON(http.StatusOK, "Account updated successfully.")
}

// Delete User Account
func DeleteUser(c echo.Context) error {
	user := new(models.User)
	deleteId := c.Param("userId")
	result := database.DB.Where("Id = ?", deleteId).Delete(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, "Account delete error.")
	}
	return c.JSON(http.StatusOK, "Account deleted successfully.")
}

// Add password for existing account