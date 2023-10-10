package api

import (
	entity "dis-test/internal/api/adapter/entity"
	model "dis-test/internal/api/app/model"
)

type DisciplineSerializer struct {
}

func NewDisciplineSerializer() *DisciplineSerializer {
	return &DisciplineSerializer{}
}

func (serializer *DisciplineSerializer) EntityArrayToDTO(disciplines []*entity.Discipline) model.DisciplineList {
	var disciplineList model.DisciplineList
	for i := 0; i < len(disciplines); i++ {
		disciplineList.Disciplines = append(disciplineList.Disciplines, serializer.EntityToDTO(disciplines[i]))
	}

	return disciplineList
}

func (serializer *DisciplineSerializer) EntityToDTO(disciplineEntity *entity.Discipline) model.Discipline {
	disciplineDTO := model.Discipline{
		ID:            disciplineEntity.ID,
		Title:         disciplineEntity.Title,
		CreditUnits:   disciplineEntity.CreditUnits,
		AcademicHours: disciplineEntity.AcademicHours,
	}

	prerequisitesList := disciplineEntity.Prerequisites

	for i := 0; i < len(prerequisitesList); i++ {
		prerequisiteDTO := model.Discipline{
			ID:            prerequisitesList[i].ID,
			Title:         prerequisitesList[i].Title,
			CreditUnits:   prerequisitesList[i].CreditUnits,
			AcademicHours: prerequisitesList[i].AcademicHours,
		}

		disciplineDTO.Prerequisites = append(disciplineDTO.Prerequisites, prerequisiteDTO)
	}

	return disciplineDTO
}
