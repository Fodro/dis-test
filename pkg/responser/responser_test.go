package pkg

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Alive bool
}

func TestResponser_SuccessfulResponse(t *testing.T) {
	responser := NewResponser()
	rr := httptest.NewRecorder()
	payload := Response{
		Alive: true,
	}

	responser.SuccessfulResponse(rr, payload)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"Alive":true}`
	assert.Equal(t, expected, rr.Body.String())

	actual := rr.Header().Get("Content-Type")
	assert.Equal(t, "application/json", actual)
}

func TestResponser_NoContentResponse(t *testing.T) {
	responser := NewResponser()
	rr := httptest.NewRecorder()

	responser.NoContentResponse(rr)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	assert.Equal(t, "null", rr.Body.String())

	actual := rr.Header().Get("Content-Type")
	assert.Equal(t, "application/json", actual)
}

func TestResponser_BadRequestResponse(t *testing.T) {
	responser := NewResponser()
	rr := httptest.NewRecorder()

	responser.BadRequestResponse(rr, errors.New("test"))

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	assert.Equal(t, `{"Code":400,"Error":"test"}`, rr.Body.String())

	actual := rr.Header().Get("Content-Type")
	assert.Equal(t, "application/json", actual)
}

func TestResponser_NotFoundResponse(t *testing.T) {
	responser := NewResponser()
	rr := httptest.NewRecorder()

	responser.NotFoundResponse(rr, errors.New("test"))

	assert.Equal(t, http.StatusNotFound, rr.Code)

	assert.Equal(t, `{"Code":404,"Error":"test"}`, rr.Body.String())

	actual := rr.Header().Get("Content-Type")
	assert.Equal(t, "application/json", actual)
}

func TestResponser_ErrorResponse(t *testing.T) {
	responser := NewResponser()
	rr := httptest.NewRecorder()

	responser.ErrorResponse(rr, http.StatusInternalServerError, errors.New("test"))

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	assert.Equal(t, `{"Code":500,"Error":"test"}`, rr.Body.String())

	actual := rr.Header().Get("Content-Type")
	assert.Equal(t, "application/json", actual)
}
