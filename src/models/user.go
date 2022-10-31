package models

import (
	"time"

	"gorm.io/gorm"
)

type UserLocation struct {
	City 				string			`json:"city" default:""`
	State				string			`json:"state" default:""`
	Country			string			`json:"country" default:""`
}

type User struct {
	ID					uint64					`gorm:"primaryKey" json:"userId"`
	CreatedAt		time.Time				`json:"-"`
	UpdatedAt		time.Time				`json:"-"`
	DeletedAt		gorm.DeletedAt	`gorm:"index,omitempty"`
	FirstName		string					`json:"firstName"`
	LastName		string					`json:"lastName"`
	Email				string					`json:"email,omitempty"`
	Password		string					`json:"password"`
	Phone				string					`json:"phone,omitempty"`
	Location 		*UserLocation  	`gorm:"embedded" json:"location"`
	Trips				[]*Trip 				`gorm:"many2many:users_trips;" json:"trips"`
}
