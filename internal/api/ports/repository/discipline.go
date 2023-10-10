package api

import (
	api "dis-test/internal/api/ports/entity"
	"gorm.io/gorm"
)

type DisciplineRepository struct {
	db *gorm.DB
}

func NewDisciplineRepository(db *gorm.DB) *DisciplineRepository {
	return &DisciplineRepository{db: db}
}

func (repo *DisciplineRepository) findAll() ([]*api.Discipline, error) {
	var disciplines []*api.Discipline

	result := repo.db.Find(&disciplines)

	return disciplines, result.Error
}

func (repo *DisciplineRepository) findById(id int) (*api.Discipline, error) {
	var discipline api.Discipline
	result := repo.db.Preload("Prerequisites").First(&discipline, id)

	return &discipline, result.Error
}

func (repo *DisciplineRepository) create(
	title string,
	creditUnits float64,
	academicHours uint32) (*api.Discipline, error) {

	discipline := api.Discipline{
		Title:         title,
		CreditUnits:   creditUnits,
		AcademicHours: academicHours,
	}

	result := repo.db.Create(&discipline)

	return &discipline, result.Error
}

func (repo *DisciplineRepository) delete(id int) error {
	result := repo.db.Delete(&api.Discipline{}, id)

	return result.Error
}

func (repo *DisciplineRepository) update(
	id int,
	title string,
	creditUnits float64,
	academicHours uint32) (*api.Discipline, error) {

	discipline, err := repo.findById(id)

	if err != nil {
		return nil, err
	}

	discipline.Title = title
	discipline.CreditUnits = creditUnits
	discipline.AcademicHours = academicHours

	result := repo.db.Save(discipline)

	return discipline, result.Error
}

func (repo *DisciplineRepository) addPrerequisite(disciplineId int, prerequisiteId int) (*api.Discipline, error) {
	discipline, err := repo.findById(disciplineId)

	if err != nil {
		return nil, err
	}

	prerequisite, err := repo.findById(prerequisiteId)

	if err != nil {
		return nil, err
	}

	err = repo.db.Model(discipline).Association("Prerequisites").Append(prerequisite)

	return discipline, err
}

func (repo *DisciplineRepository) deletePrerequisite(disciplineId int, prerequisiteId int) (*api.Discipline, error) {
	discipline, err := repo.findById(disciplineId)

	if err != nil {
		return nil, err
	}

	prerequisite, err := repo.findById(prerequisiteId)

	if err != nil {
		return nil, err
	}

	err = repo.db.Model(discipline).Association("Prerequisites").Delete(prerequisite)

	return discipline, err
}
