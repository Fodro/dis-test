package api

import (
	repository "dis-test/internal/api/adapter/repository"
	model "dis-test/internal/api/app/model"
	dto "dis-test/internal/api/app/serializer"
)

type DisciplineService struct {
	repo       *repository.Repo
	serializer *dto.Serializer
}

func NewDisciplineService(repo *repository.Repo, serializer *dto.Serializer) *DisciplineService {
	return &DisciplineService{repo, serializer}
}

func (service *DisciplineService) Create(disciplineDTO *model.Discipline) (*model.Discipline, error) {
	disciplineEntity, err := service.repo.Create(disciplineDTO.Title, disciplineDTO.CreditUnits, disciplineDTO.AcademicHours)

	if err != nil {
		return nil, err
	}

	result := service.serializer.EntityToDTO(disciplineEntity)

	return &result, nil
}

func (service *DisciplineService) GetList() (*model.DisciplineList, error) {
	disciplineEntityList, err := service.repo.FindAll()

	if err != nil {
		return nil, err
	}

	disciplineDTOList := service.serializer.EntityArrayToDTO(disciplineEntityList)

	return &disciplineDTOList, nil
}

func (service *DisciplineService) GetById(id int) (*model.Discipline, error) {
	disciplineEntity, err := service.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	disciplineDTO := service.serializer.EntityToDTO(disciplineEntity)

	return &disciplineDTO, nil
}
