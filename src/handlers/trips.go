package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Find All Trips
func FindAllTrips(c echo.Context) error {
	var trips []models.Trip

	result := database.DB.Preload("Team", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "FirstName", "LastName")
	}).Find(&trips)

	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"message": "Unable to locate trips.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"trips": trips,
	})
}

// Find Single Trip
func FindTrip(c echo.Context) error {
	trip := new(models.Trip)
	tripId := c.Param("tripId")
	result := database.DB.Preload("Team", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "FirstName", "LastName")
	}).First(&trip, tripId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Could not find trip details.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"trip": trip,
	})
}

// Create a New Trip
func CreateTrip(c echo.Context) error {
	newTrip := new(models.Trip)
	if err := c.Bind(&newTrip); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"message": "Unable to bind data.",
		})
	}
	result := database.DB.Create(&newTrip)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Unable to create new trip.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "Trip added successfully.",
	})
}

// Update a Trip
func UpdateTrip(c echo.Context) error {
	trip := new(models.Trip)
	updateId := c.Param("tripId")
	result := database.DB.First(&trip, updateId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Could not find trip in database.",
		})
	}
	c.Bind(&trip)
	result2 := database.DB.Save(&trip)
	if result2.Error != nil {
		log.Fatal(result2.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Trip update error.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "Trip updated successfully.",
	})
}

// Delete a Trip
func DeleteTrip(c echo.Context) error {
	trip := new(models.Trip)
	deleteId := c.Param("tripId")
	result := database.DB.Where("Id = ?", deleteId).Delete(&trip)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Delete trip error.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "Trip deleted successfully.",
	})
}

// Add a User to a Trip's Team
func AddTeamMember(c echo.Context) error {
	user := new(models.User)
	userId := c.Param("userId")

	tripId, err := strconv.ParseUint(c.Param("tripId"), 10, 32)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"message": "Invalid tripId.",
		})
	}

	result := database.DB.Find(&user, userId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"message": "Invalid userId.",
		})
	}

	err2 := database.DB.Model(&user).Association("Trips").Append(&models.Trip{ID: tripId})
	if err2 != nil {
		log.Fatal(err2)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Association error.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "User added to trip successfully.",
	})
}

// Remove a User from a Trip's Team
func RemoveTeamMember(c echo.Context) error {
	user := new(models.User)
	userId := c.Param("userId")

	tripId, err := strconv.ParseUint(c.Param("tripId"), 10, 32)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"message": "Invalid tripId.",
		})
	}

	result := database.DB.Find(&user, userId)
	if result.Error != nil {
		log.Fatal(result.Error)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": 400,
			"message": "Invalid userId.",
		})
	}

	err2 := database.DB.Model(&user).Association("Trips").Delete(&models.Trip{ID: tripId})
	if err2 != nil {
		log.Fatal(err2)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": "Association error.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"message": "User removed from trip successfully.",
	})
}
