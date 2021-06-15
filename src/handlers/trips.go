package handlers

import (
	"log"
	"net/http"

	"github.com/kinr-jay/hee-haw-go/src/database"
	"github.com/kinr-jay/hee-haw-go/src/models"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func CreateTrip(c echo.Context) error {

		trip := models.Trip{
			Title: "Sample Trip",
			Description: "Lorem ipsum dolor yippiekayay mf",
			Image: "https://nothingisreal.net",
			StartDate: "06/14/2021",
			EndDate: "06/14/2021",
			Area: "The backyard",
			Regs: "Standard WA / NF regs, bear cannisters not required",
			Muster: "The Ace Hardware Store on Colfax",
			Distance: 5,
			Elevation: 1500,
			GroupSize: 10,
			GearList: pq.StringArray([]string{"boots", "backpack", "tent"}),
			Completed: false,
			Report: "",
			Team: []*models.User{},
		}

		result := database.DB.Create(&trip)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		return c.JSON(http.StatusOK, "Trip added successfully.")
	}

	func FindAllTrips(c echo.Context) error {
		var trips []models.Trip
		database.DB.Find(&trips)
		return c.JSON(http.StatusOK, trips)
	}

	func Test(c echo.Context) error {
		newTrip := new(models.Trip)
		c.Bind(&newTrip)
		result := database.DB.Create(&newTrip)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		return c.JSON(http.StatusOK, "Trip added successfully.")
	}