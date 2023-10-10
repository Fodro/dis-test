package api

import (
	entity "dis-test/internal/api/adapter/entity"
	model "dis-test/internal/api/app/model"
)

type Serializer struct {
	Discipline
}

type Discipline interface {
	EntityArrayToDTO(disciplines []*entity.Discipline) model.DisciplineList
	EntityToDTO(disciplineEntity *entity.Discipline) model.Discipline
}

func NewSerializer() *Serializer {
	return &Serializer{
		Discipline: NewDisciplineSerializer(),
	}
}
