package handlers

import (
	"log"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"
	"github.com/labstack/echo"
)

// func CreateTrip(c echo.Context) error {

// 	trip := models.Trip{
// 		Title: "Sample Trip",
// 		Description: "Lorem ipsum dolor yippiekayay mf",
// 		Image: "https://nothingisreal.net",
// 		StartDate: "06/14/2021",
// 		EndDate: "06/14/2021",
// 		Area: "The backyard",
// 		Regs: "Standard WA / NF regs, bear cannisters not required",
// 		Muster: "The Ace Hardware Store on Colfax",
// 		Distance: 5,
// 		Elevation: 1500,
// 		GroupSize: 10,
// 		GearList: pq.StringArray([]string{"boots", "backpack", "tent"}),
// 		Completed: false,
// 		Report: "",
// 		Team: []*models.User{},
// 	}

// 	result := database.DB.Create(&trip)
// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}
// 	return c.JSON(http.StatusOK, "Trip added successfully.")
// }

func FindAllTrips(c echo.Context) error {
	var trips []models.Trip
	database.DB.Find(&trips)
	return c.JSON(http.StatusOK, trips)
}

func CreateTrip(c echo.Context) error {
	newTrip := new(models.Trip)
	if err := c.Bind(&newTrip); err != nil {
		log.Fatal(err)
	}
	result := database.DB.Create(&newTrip)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return c.JSON(http.StatusOK, "Trip added successfully.")
}

func UpdateTrip(c echo.Context) error {
	trip := new(models.Trip)
	updateId := c.Param("tripId")
	result := database.DB.First(&trip, updateId)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.Bind(&trip)
	result2 := database.DB.Save(&trip)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	return c.JSON(http.StatusOK, "Trip updated successfully.")
}

func DeleteTrip(c echo.Context) error {
	var trip models.Trip
	deleteId := c.Param("tripId")
	result := database.DB.Where("Id = ?", deleteId).Delete(&trip)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return c.JSON(http.StatusOK, "Trip deleted successfully.")
}