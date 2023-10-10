package api

import (
	model "dis-test/internal/api/app/model"
	app "dis-test/internal/api/app/service"
	pkg "dis-test/pkg/responser"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DisciplineHandler struct {
	r         *mux.Router
	service   *app.Service
	responser *pkg.Responser
}

func NewDisciplineHandler(r *mux.Router, service *app.Service, responser *pkg.Responser) *DisciplineHandler {
	disciplineHandler := DisciplineHandler{r, service, responser}
	disciplineHandler.addRoutes()
	return &disciplineHandler
}

func (handler *DisciplineHandler) addRoutes() {
	disciplines := handler.r.PathPrefix("/discipline").Subrouter()

	disciplines.HandleFunc("/", handler.HandleList).Methods("GET")
	disciplines.HandleFunc("/", handler.HandleCreate).Methods("POST")
	disciplines.HandleFunc("/{id}", handler.HandleShow).Methods("GET")
	disciplines.HandleFunc("/{id}", handler.HandleUpdate).Methods("PUT")
	disciplines.HandleFunc("/{id}", handler.HandleDelete).Methods("DELETE")
	disciplines.HandleFunc("/{id}/prerequisite", handler.HandleAddPrerequisite).Methods("POST")
	disciplines.HandleFunc("/{id}/prerequisite", handler.HandleDeletePrerequisite).Methods("DELETE")
}

func (handler *DisciplineHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	disciplineList, err := handler.service.Discipline.GetList()

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	handler.responser.SuccessfulResponse(w, disciplineList)
}

func (handler *DisciplineHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	discipline, err := handler.service.Discipline.GetById(id)

	if err != nil {
		handler.responser.NotFoundResponse(w, err)
		return
	}

	handler.responser.SuccessfulResponse(w, discipline)
}

func (handler *DisciplineHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var newDiscipline model.Discipline

	err := json.NewDecoder(r.Body).Decode(&newDiscipline)

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	result, err := handler.service.Discipline.Create(&newDiscipline)

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	handler.responser.SuccessfulResponse(w, result)
}

func (handler *DisciplineHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var newDiscipline model.Discipline

	err := json.NewDecoder(r.Body).Decode(&newDiscipline)

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	result, err := handler.service.Discipline.UpdateById(id, &newDiscipline)

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	handler.responser.SuccessfulResponse(w, result)
}

func (handler *DisciplineHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := handler.service.Discipline.DeleteById(id)

	if err != nil {
		handler.responser.NotFoundResponse(w, err)
		return
	}

	handler.responser.NoContentResponse(w)
}

func (handler *DisciplineHandler) HandleAddPrerequisite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var prerequisite model.Prerequisite

	err := json.NewDecoder(r.Body).Decode(&prerequisite)

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	discipline, err := handler.service.Discipline.AddPrerequisiteById(id, prerequisite.ID)

	if err != nil {
		handler.responser.NotFoundResponse(w, err)
		return
	}

	handler.responser.SuccessfulResponse(w, discipline)
}

func (handler *DisciplineHandler) HandleDeletePrerequisite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var prerequisite model.Prerequisite

	err := json.NewDecoder(r.Body).Decode(&prerequisite)

	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}

	discipline, err := handler.service.Discipline.DeletePrerequisiteById(id, prerequisite.ID)

	if err != nil {
		handler.responser.NotFoundResponse(w, err)
		return
	}

	handler.responser.SuccessfulResponse(w, discipline)
}
