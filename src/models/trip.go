package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Trip struct {
	ID					uint64					`gorm:"primaryKey" json:"tripId"`
	CreatedAt		time.Time				`json:"-"`
	UpdatedAt		time.Time				`json:"-"`
	DeletedAt		gorm.DeletedAt	`gorm:"index" json:"-"`
	Title				string					`json:"title,omitempty"`
	Description string					`json:"description,omitempty"`
	Image				string					`json:"image,omitempty"`
	StartDate		string 					`json:"startDate,omitempty"`
	EndDate			string	 				`json:"endDate,omitempty"`
	Area				string					`json:"area,omitempty"`
	Regs				string					`json:"regs,omitempty"`
	Muster			string					`json:"muster,omitempty"`
	Distance		uint8 					`json:"distance,omitempty"`
	Elevation 	uint16					`json:"elevation,omitempty"`
	GroupSize		uint8 					`json:"groupSize,omitempty"`
	GearList		pq.StringArray	`gorm:"type:text[]" json:"gearList,omitempty"`
	Team				[]*User 				`gorm:"many2many:users_trips" json:"team"`
	Completed		bool  					`json:"completed,omitempty"`
	Report			string					`json:"report,omitempty"`
}
