package api

import (
	repository "dis-test/internal/api/adapter/repository"
	dto "dis-test/internal/api/app/serializer"
)

type Service struct {
	Discipline
}

type Discipline interface {
}

func NewService(repo *repository.Repo, serializer *dto.Serializer) *Service {
	return &Service{
		Discipline: NewDisciplineService(repo, serializer),
	}
}
