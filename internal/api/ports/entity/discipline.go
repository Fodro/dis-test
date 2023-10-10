package api

import "gorm.io/gorm"

type Discipline struct {
	gorm.Model
	Title         string
	CreditUnits   float64
	AcademicHours uint32
	Prerequisites []Discipline `gorm:"many2many:discipline_prerequisites"`
}
