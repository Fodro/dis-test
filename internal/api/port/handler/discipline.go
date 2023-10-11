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
	disciplines := handler.r.PathPrefix("/api/discipline").Subrouter()

	disciplines.HandleFunc("/", handler.HandleList).Methods(http.MethodGet)
	disciplines.HandleFunc("/", handler.HandleCreate).Methods(http.MethodPost)
	disciplines.HandleFunc("/{id}", handler.HandleShow).Methods(http.MethodGet)
	disciplines.HandleFunc("/{id}", handler.HandleUpdate).Methods(http.MethodPut)
	disciplines.HandleFunc("/{id}", handler.HandleDelete).Methods(http.MethodDelete)
	disciplines.HandleFunc("/{id}/prerequisite", handler.HandleAddPrerequisite).Methods(http.MethodPatch)
	disciplines.HandleFunc("/{id}/prerequisite", handler.HandleDeletePrerequisite).Methods(http.MethodDelete)
}

// HandleList godoc
// @Summary      List all disciplines
// @Description  List of all disciplines existing in DB
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.DisciplineList
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/ [get]
func (handler *DisciplineHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	disciplineList, err := handler.service.Discipline.GetList()
	if err != nil {
		handler.responser.BadRequestResponse(w, err)
		return
	}
	handler.responser.SuccessfulResponse(w, disciplineList)
}

// HandleShow godoc
// @Summary      Show discipline
// @Description  Find and serve concrete discipline by id
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Discipline ID"
// @Success      200  {object}  model.Discipline
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/{id} [get]
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

// HandleCreate godoc
// @Summary      Create discipline
// @Description  Create new discipline by constant data
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param data body model.DisciplineInput true "The input discipline struct"
// @Success      200  {object}  model.Discipline
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/ [post]
func (handler *DisciplineHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var newDiscipline model.DisciplineInput

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

// HandleUpdate godoc
// @Summary      Update discipline
// @Description  Find and update concrete discipline by id
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Discipline ID"
// @Param data body model.DisciplineInput true "The input discipline struct"
// @Success      200  {object}  model.Discipline
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/{id} [put]
func (handler *DisciplineHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var newDiscipline model.DisciplineInput

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

// HandleDelete godoc
// @Summary      Delete discipline
// @Description  Find and delete concrete discipline by id
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Discipline ID"
// @Success      204
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/{id} [delete]
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

// HandleAddPrerequisite godoc
// @Summary      Add prerequisite
// @Description  Find discipline and prerequisite, link them
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Discipline ID"
// @Param data body model.Prerequisite true "The input struct contains discipline id to add as prerequisite"
// @Success      200  {object}  model.Discipline
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/{id}/prerequisite [patch]
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

// HandleDeletePrerequisite godoc
// @Summary      Delete prerequisite
// @Description  Find discipline and delete relation to prerequisite by ids
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Discipline ID"
// @Param data body model.Prerequisite true "The input struct contains id of prerequisite to delete"
// @Success      200  {object}  model.Discipline
// @Failure      400  {object}  pkg.Error
// @Failure      404  {object}  pkg.Error
// @Failure      500  {object}  pkg.Error
// @Router       /discipline/{id}/prerequisite [delete]
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
