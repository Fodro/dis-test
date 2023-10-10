package api

import (
	app "dis-test/internal/api/app/service"
	pkg "dis-test/pkg/responser"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Discipline
}

type Discipline interface {
	addRoutes()
	HandleList(w http.ResponseWriter, r *http.Request)
	HandleShow(w http.ResponseWriter, r *http.Request)
	HandleCreate(w http.ResponseWriter, r *http.Request)
	HandleUpdate(w http.ResponseWriter, r *http.Request)
	HandleDelete(w http.ResponseWriter, r *http.Request)
	HandleAddPrerequisite(w http.ResponseWriter, r *http.Request)
	HandleDeletePrerequisite(w http.ResponseWriter, r *http.Request)
}

func NewHandler(r *mux.Router, service *app.Service, responser *pkg.Responser) *Handler {
	return &Handler{
		Discipline: NewDisciplineHandler(r, service, responser),
	}
}
