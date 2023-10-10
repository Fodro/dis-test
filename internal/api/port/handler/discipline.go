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

	disciplines.HandleFunc("/", handler.HandleList)
	disciplines.HandleFunc("/{id}", handler.HandleShow)
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
