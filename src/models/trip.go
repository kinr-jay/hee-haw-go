package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	Title				string					`json:"title"`
	Description string					`json:"description"`
	Image				string					`json:"image"`
	StartDate		string 					`json:"startDate"`
	EndDate			string	 				`json:"endDate"`
	Area				string					`json:"area"`
	Regs				string					`json:"regs"`
	Muster			string					`json:"muster"`
	Distance		uint8 					`json:"distance"`
	Elevation 	uint16					`json:"elevation"`
	GroupSize		uint8 					`json:"groupSize"`
	GearList		pq.StringArray	`gorm:"type:text[];" json:"gearList"`
	Team				[]*User 				`gorm:"many2many:users_trips;"`
	Completed		bool  					`json:"completed"`
	Report			string					`json:"report"`
}
