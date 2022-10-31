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
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Unable to create new account.",
		})
	}
	
	token, err2 := CreateJWT(user.Email, user.ID)
	if err2 != nil {
		fmt.Println("Error while creating JWT token: ", err2)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Unable to login.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
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
	}).Select("id", "first_name", "last_name", "email", "phone", "city", "state", "country").First(&user, userId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Could not find user account.",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// Update User Account
func UpdateUser(c echo.Context) error {
	updatedUser := new(models.User)
	databaseUser := new(models.User)
	updateId := c.Param("userId")

	c.Bind(&updatedUser)

	result := database.DB.First(&databaseUser, updateId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Could not find user account.",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(updatedUser.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status": 401,
			"message": "Incorrect password provided.",
		})
	}
	updatedUser.Password = databaseUser.Password

	result2 := database.DB.Save(&updatedUser)
	if result2.Error != nil {
		log.Fatal(result2.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Account update error.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "User account updated successfully",
	})

}

type verifyDelete struct {
	Password string `json:"password"`
}

// Delete User Account
func DeleteUser(c echo.Context) error {
	user := new(models.User)
	deleteId := c.Param("userId")

	verifyDelete := new(verifyDelete)
	c.Bind(&verifyDelete)

	result := database.DB.First(&user, deleteId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Could not find user account.",
		})
	}
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(verifyDelete.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status": 401,
			"message": "Incorrect password provided.",
		})
	}

	result2 := database.DB.Where("Id = ?", deleteId).Delete(&user)
	if result2.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Account delete error.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "Account deleted successfully.",
	})
}
