package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email				string 		`json:"email"`
	Password		string		`json:"password"`
}

///////// JWT Functions ////////////////////
//////// Create JWT
func CreateJWT(email string, userId uint64) (string, error) {
	
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id": userId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := rawToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

// User Login Function
func Login(c echo.Context) error {
	user := new(models.User)
	login := new(login)
	c.Bind(&login)
	
	result := database.DB.Where("email = ?", login.Email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status": 401,
			"message": "There is no account registered under that email.",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status": 401,
			"message": "Incorrect password.",
		})
	}

	token, err2 := CreateJWT(user.Email, user.ID)
	if err2 != nil {
		fmt.Println("Error while creating JWT token: ", err2)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Unable to generate JWT token.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "Login successful.",
		"userId": user.ID,
		"token": token,
	})
}
