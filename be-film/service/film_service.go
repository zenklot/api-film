package service

import "github.com/zenklot/api-film/model/web"

type FilmService interface {
	Create(request web.FilmCreateRequest) web.FilmResponse
	Delete(filmId string)
	FindById(filmId string) web.FilmResponse
	FindAll(page int) ([]web.FilmResponse, int)
	FindByTitle(page int, title string) ([]web.FilmResponse, int)
}
