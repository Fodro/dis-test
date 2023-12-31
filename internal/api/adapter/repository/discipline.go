package api

import (
	entity "dis-test/internal/api/adapter/entity"
	"gorm.io/gorm"
)

type DisciplineRepository struct {
	db *gorm.DB
}

func NewDisciplineRepository(db *gorm.DB) *DisciplineRepository {
	return &DisciplineRepository{db: db}
}

func (repo *DisciplineRepository) FindAll() ([]*entity.Discipline, error) {
	var disciplines []*entity.Discipline

	result := repo.db.Preload("Prerequisites").Find(&disciplines)

	return disciplines, result.Error
}

func (repo *DisciplineRepository) FindById(id int) (*entity.Discipline, error) {
	var discipline entity.Discipline
	result := repo.db.Preload("Prerequisites").First(&discipline, id)

	return &discipline, result.Error
}

func (repo *DisciplineRepository) Create(title string, creditUnits float32, academicHours uint32) (*entity.Discipline, error) {

	discipline := entity.Discipline{
		Title:         title,
		CreditUnits:   creditUnits,
		AcademicHours: academicHours,
	}

	result := repo.db.Create(&discipline)

	return &discipline, result.Error
}

func (repo *DisciplineRepository) Delete(id int) error {
	result := repo.db.Delete(&entity.Discipline{}, id)

	return result.Error
}

func (repo *DisciplineRepository) Update(id int, title string, creditUnits float32, academicHours uint32) (*entity.Discipline, error) {
	var discipline entity.Discipline
	result := repo.db.First(&discipline, id)

	if result.Error != nil {
		return nil, result.Error
	}

	discipline.Title = title
	discipline.CreditUnits = creditUnits
	discipline.AcademicHours = academicHours

	result = repo.db.Save(discipline)

	return &discipline, result.Error
}

func (repo *DisciplineRepository) AddPrerequisite(disciplineId int, prerequisiteId int) (*entity.Discipline, error) {
	discipline, err := repo.FindById(disciplineId)

	if err != nil {
		return nil, err
	}

	prerequisite, err := repo.FindById(prerequisiteId)

	if err != nil {
		return nil, err
	}

	err = repo.db.Model(discipline).Association("Prerequisites").Append(prerequisite)

	return discipline, err
}

func (repo *DisciplineRepository) DeletePrerequisite(disciplineId int, prerequisiteId int) (*entity.Discipline, error) {
	discipline, err := repo.FindById(disciplineId)

	if err != nil {
		return nil, err
	}

	prerequisite, err := repo.FindById(prerequisiteId)

	if err != nil {
		return nil, err
	}

	err = repo.db.Model(discipline).Association("Prerequisites").Delete(prerequisite)

	return discipline, err
}
