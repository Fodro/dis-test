package api

import (
	repository "dis-test/internal/api/adapter/repository"
	model "dis-test/internal/api/app/model"
	dto "dis-test/internal/api/app/serializer"
)

type Service struct {
	Discipline
}

type Discipline interface {
	Create(disciplineDTO *model.Discipline) (*model.Discipline, error)
	GetList() (*model.DisciplineList, error)
	GetById(id int) (*model.Discipline, error)
	DeleteById(id int) error
	UpdateById(id int, disciplineDTO *model.Discipline) (*model.Discipline, error)
	AddPrerequisiteById(disciplineId int, prerequisiteId int) (*model.Discipline, error)
	DeletePrerequisiteById(disciplineId int, prerequisiteId int) (*model.Discipline, error)
}

func NewService(repo *repository.Repo, serializer *dto.Serializer) *Service {
	return &Service{
		Discipline: NewDisciplineService(repo, serializer),
	}
}
