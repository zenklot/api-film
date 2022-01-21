package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FilmController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByTitle(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
