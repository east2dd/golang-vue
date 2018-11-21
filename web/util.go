package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pasperry/govolie/util/log"
)

var (
	ErrParamNotFound       = errors.New("Param could not be found.")
	ErrParamCouldNotAssign = errors.New("Param could not be assigned.")
)

type RenderActionRequestFunc func(http.ResponseWriter, *http.Request) RenderAction

type RenderAction interface {
	Render(http.ResponseWriter)
}

type ErrorAction struct {
	Error error
}

func (action ErrorAction) Render(c http.ResponseWriter) {
	if action.Error != nil {
		log.Error.Println(action.Error)
	}

	c.WriteHeader(http.StatusInternalServerError)
	c.Write([]byte("An error has occured."))
}

type NotFoundAction struct{}

func (action NotFoundAction) Render(c http.ResponseWriter) {
	c.WriteHeader(http.StatusNotFound)
	c.Write([]byte("Document not found."))
}

type JsonAction struct {
	Data interface{}
}

func (action JsonAction) Render(c http.ResponseWriter) {
	b, err := json.Marshal(&action.Data)
	if err != nil {
		panic(err)
		return
	}

	c.WriteHeader(http.StatusOK)
	c.Write([]byte(b))
}

func HandleAction(actionFunc RenderActionRequestFunc) func(http.ResponseWriter, *http.Request) {
	return func(c http.ResponseWriter, r *http.Request) {
		result := actionFunc(c, r)
		result.Render(c)
	}
}
