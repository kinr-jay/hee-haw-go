package models

import (
	"gorm.io/gorm"
)

type UserLocation struct {
	gorm.Model
	City 				string
	State				string
	Country			string
}

type User struct {
	gorm.Model
	FirstName		string
	LastName		string
	Email				string
	Phone				string
	Location 		UserLocation `gorm:"embedded"`
	Trips				[]*Trip `gorm:"many2many:users_trips;"`
}
