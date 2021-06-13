package models

import (
	"time"
)

type Trip struct {
	ID					uint
	Title				string
	StartDate		time.Time
	EndDate			time.Time
	Area				string
	Team				[]*User `gorm:"many2many:users_trips;"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		time.Time
}
