package api

import (
	model "dis-test/internal/api/app/model"
	appService "dis-test/internal/api/app/service"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strconv"
)

const (
	createMessage             = "create"
	deleteMessage             = "delete"
	updateMessage             = "update"
	addPrerequisiteMessage    = "addPrerequisite"
	deletePrerequisiteMessage = "deletePrerequisite"
)

type DisciplineMessageHandler struct {
	service *appService.Service
}

func NewDisciplineMessageHandler(service *appService.Service) *DisciplineMessageHandler {
	return &DisciplineMessageHandler{service}
}

func (handler *DisciplineMessageHandler) HandleEvent(ev *kafka.Message) {
	if string(ev.Key) == createMessage {
		handler.handleCreate(ev.Value)
		return
	}
	if string(ev.Key) == deleteMessage {
		handler.handleDelete(ev.Value)
		return
	}
	if string(ev.Key) == updateMessage {
		handler.handleUpdate(ev.Value)
		return
	}
	if string(ev.Key) == addPrerequisiteMessage {
		handler.handleAddPrerequisite(ev.Value)
		return
	}
	if string(ev.Key) == deletePrerequisiteMessage {
		handler.handleDeletePrerequisite(ev.Value)
		return
	}
}

func (handler *DisciplineMessageHandler) handleCreate(payload []byte) {
	var discipline model.DisciplineInput

	json.Unmarshal(payload, &discipline)

	handler.service.Discipline.Create(&discipline)
}

func (handler *DisciplineMessageHandler) handleDelete(payload []byte) {
	id, _ := strconv.Atoi(string(payload))

	handler.service.Discipline.DeleteById(id)
}

func (handler *DisciplineMessageHandler) handleUpdate(payload []byte) {
	var discipline model.Discipline

	json.Unmarshal(payload, &discipline)

	handler.service.Discipline.UpdateById(int(discipline.ID), &model.DisciplineInput{Title: discipline.Title, AcademicHours: discipline.AcademicHours, CreditUnits: discipline.CreditUnits})
}

func (handler *DisciplineMessageHandler) handleAddPrerequisite(payload []byte) {
	var prerequisite model.PrerequisiteFull

	json.Unmarshal(payload, &prerequisite)

	handler.service.Discipline.AddPrerequisiteById(prerequisite.DisciplineID, prerequisite.PrerequisiteID)
}

func (handler *DisciplineMessageHandler) handleDeletePrerequisite(payload []byte) {
	var prerequisite model.PrerequisiteFull

	json.Unmarshal(payload, &prerequisite)

	handler.service.Discipline.DeletePrerequisiteById(prerequisite.DisciplineID, prerequisite.PrerequisiteID)
}
