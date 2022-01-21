package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zenklot/api-film/controller"
	"github.com/zenklot/api-film/exception"
)

func NewRouter(filmController controller.FilmController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/film", filmController.Create)
	router.GET("/api/films", filmController.FindAll)
	router.GET("/api/films/search/:title", filmController.FindByTitle)
	router.GET("/api/film/:filmId", filmController.FindById)
	router.DELETE("/api/film/:filmId", filmController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
