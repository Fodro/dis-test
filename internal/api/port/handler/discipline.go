package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type DisciplineHandler struct {
	r *mux.Router
}

func NewDisciplineHandler(r *mux.Router) *DisciplineHandler {
	disciplineHandler := DisciplineHandler{r}
	disciplineHandler.addRoutes()
	return &disciplineHandler
}

func (handler *DisciplineHandler) addRoutes() {
	disciplines := handler.r.PathPrefix("/discipline").Subrouter()

	disciplines.HandleFunc("/", handler.HandleList).Methods("GET")
	disciplines.HandleFunc("/{id}", handler.HandleShow)
	disciplines.HandleFunc("/", handler.HandleCreate).Methods("POST")
	disciplines.HandleFunc("/{id}", handler.HandleUpdate).Methods("PUT")
	disciplines.HandleFunc("/{id}", handler.HandleDelete).Methods("DELETE")
	disciplines.HandleFunc("/{id}/prerequisite", handler.HandleAddPrerequisite).Methods("POST")
	disciplines.HandleFunc("/{id}/prerequisite", handler.HandleDeletePrerequisite).Methods("DELETE")
}

func (handler *DisciplineHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func (handler *DisciplineHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("%s", id)
	fmt.Fprintf(w, response)
}

func (handler *DisciplineHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func (handler *DisciplineHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func (handler *DisciplineHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func (handler *DisciplineHandler) HandleAddPrerequisite(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func (handler *DisciplineHandler) HandleDeletePrerequisite(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
