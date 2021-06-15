package models

import (
	"gorm.io/gorm"
)

type UserLocation struct {
	gorm.Model
	City 				string			`json:"city"`
	State				string			`json:"state"`
	Country			string			`json:"country"`
}

type User struct {
	gorm.Model
	FirstName		string				`json:"firstName"`
	LastName		string				`json:"lastName"`
	Email				string				`json:"email"`
	Password		string				`json:"password"`
	Phone				string				`json:"phone"`
	Location 		UserLocation  `gorm:"embedded" json:"location"`
	Trips				[]*Trip 			`gorm:"many2many:users_trips;"`
}
