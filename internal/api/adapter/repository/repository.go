package api

import (
	api "dis-test/internal/api/adapter/entity"
	"gorm.io/gorm"
)

type Repo struct {
	Discipline
}

type Discipline interface {
	FindAll() ([]*api.Discipline, error)
	FindById(id int) (*api.Discipline, error)
	Create(title string, creditUnits float64, academicHours uint32) (*api.Discipline, error)
	Delete(id int) error
	Update(id int, title string, creditUnits float64, academicHours uint32) (*api.Discipline, error)
	AddPrerequisite(disciplineId int, prerequisiteId int) (*api.Discipline, error)
	DeletePrerequisite(disciplineId int, prerequisiteId int) (*api.Discipline, error)
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		Discipline: NewDisciplineRepository(db),
	}
}
