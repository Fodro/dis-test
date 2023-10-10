package pkg

import (
	"encoding/json"
	"net/http"
)

type Responser struct{}

func NewResponser() *Responser {
	return &Responser{}
}

func (responser *Responser) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (responser *Responser) SuccessfulResponse(w http.ResponseWriter, payload interface{}) {
	responser.respondWithJSON(w, http.StatusOK, payload)
}

func (responser *Responser) ErrorResponse(w http.ResponseWriter, code int, err error) {
	errorMessage := Error{
		Code:  code,
		Error: err.Error(),
	}

	responser.respondWithJSON(w, code, errorMessage)
}

func (responser *Responser) NotFoundResponse(w http.ResponseWriter, err error) {
	responser.ErrorResponse(w, http.StatusNotFound, err)
}

func (responser *Responser) BadRequestResponse(w http.ResponseWriter, err error) {
	responser.ErrorResponse(w, http.StatusBadRequest, err)
}

func (responser *Responser) NoContentResponse(w http.ResponseWriter) {
	responser.respondWithJSON(w, http.StatusNoContent, nil)
}
