package api

import (
	entity "dis-test/internal/api/adapter/entity"
	model "dis-test/internal/api/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisciplineSerializer_EntityToDTO(t *testing.T) {
	serializer := NewSerializer()
	disciplineEntity := entity.Discipline{
		Title:         "TEST",
		CreditUnits:   1.5,
		AcademicHours: 2,
		Prerequisites: []entity.Discipline{{
			Title:         "PREREQUISITE",
			CreditUnits:   6.7,
			AcademicHours: 9,
		}},
	}
	disciplineEntity.ID = 1

	got := serializer.EntityToDTO(&disciplineEntity)

	expected := model.Discipline{
		ID:            1,
		Title:         "TEST",
		CreditUnits:   1.5,
		AcademicHours: 2,
		Prerequisites: []model.Discipline{{
			Title:         "PREREQUISITE",
			CreditUnits:   6.7,
			AcademicHours: 9,
		}},
	}

	assert.Equal(t, expected, got)
}

func TestDisciplineSerializer_EntityArrayToDTO(t *testing.T) {
	serializer := NewSerializer()
	disciplineEntity := entity.Discipline{
		Title:         "TEST",
		CreditUnits:   1.5,
		AcademicHours: 2,
		Prerequisites: []entity.Discipline{{
			Title:         "PREREQUISITE",
			CreditUnits:   6.7,
			AcademicHours: 9,
		}},
	}
	disciplineEntity.ID = 1

	got := serializer.EntityArrayToDTO([]*entity.Discipline{&disciplineEntity})

	expected := model.DisciplineList{Disciplines: []model.Discipline{{
		ID:            1,
		Title:         "TEST",
		CreditUnits:   1.5,
		AcademicHours: 2,
		Prerequisites: []model.Discipline{{
			Title:         "PREREQUISITE",
			CreditUnits:   6.7,
			AcademicHours: 9,
		}},
	}}}

	assert.Equal(t, expected, got)
}
