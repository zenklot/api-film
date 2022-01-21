package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/api-film/model/web"
	"github.com/zenklot/api-film/service"
)

type FilmControllerImpl struct {
	FilmService service.FilmService
}

func NewFilmController(filmService service.FilmService) FilmController {
	return &FilmControllerImpl{
		FilmService: filmService,
	}

}

func (controller *FilmControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var page int
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if page == 0 || err != nil {
		page = 1
	}
	filmResponse, totalPage := controller.FilmService.FindAll(page)
	webResponse := web.WebResponsePaginate{
		Code:        200,
		Status:      "OK",
		Data:        filmResponse,
		TotalPage:   totalPage,
		CurrentPage: page,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *FilmControllerImpl) FindByTitle(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var page int
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if page == 0 || err != nil {
		page = 1
	}
	title := params.ByName("title")
	filmResponse, totalPage := controller.FilmService.FindByTitle(page, title)
	webResponse := web.WebResponsePaginate{
		Code:        200,
		Status:      "OK",
		Data:        filmResponse,
		TotalPage:   totalPage,
		CurrentPage: page,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *FilmControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	filmId := params.ByName("filmId")
	filmResponse := controller.FilmService.FindById(filmId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   filmResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *FilmControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	filmCreateRequest := web.FilmCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&filmCreateRequest)
	if err != nil {
		panic(err)
	}

	filmResponse := controller.FilmService.Create(filmCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   filmResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}

}

func (controller *FilmControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	filmId := params.ByName("filmId")
	controller.FilmService.Delete(filmId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}
