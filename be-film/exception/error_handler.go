package exception

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zenklot/api-film/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}
	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "*")
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		encoder := json.NewEncoder(writer)
		err := encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}

		return true
	} else {
		return false
	}
}
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "*")
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		encoder := json.NewEncoder(writer)
		err := encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Methods", "*")
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}
