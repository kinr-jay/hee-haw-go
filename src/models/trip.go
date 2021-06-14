package models

import (
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	Title				string
	StartDate		time.Time
	EndDate			time.Time
	Area				string
	Team				[]*User `gorm:"many2many:users_trips;"`
}
