package service

import (
	"github.com/go-playground/validator/v10"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/zenklot/api-film/model/domain"
	"github.com/zenklot/api-film/model/web"
	"github.com/zenklot/api-film/repository"
)

type FilmServiceImpl struct {
	FilmRepository repository.FilmRepository
	DB             *scribble.Driver
	Validate       *validator.Validate
}

func NewFilmService(filmRepository repository.FilmRepository, DB *scribble.Driver, validate *validator.Validate) FilmService {
	return &FilmServiceImpl{
		FilmRepository: filmRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *FilmServiceImpl) Create(request web.FilmCreateRequest) web.FilmResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	db := service.DB
	film := domain.Film{
		Title:    request.Title,
		Genre:    request.Genre,
		Rating:   request.Rating,
		Duration: request.Duration,
		Quality:  request.Quality,
		Trailer:  request.Trailer,
		Watch:    request.Watch,
	}
	filmSaved := service.FilmRepository.Save(db, film)
	return web.FilmResponse{
		Id:       filmSaved.Id,
		Title:    filmSaved.Title,
		Genre:    filmSaved.Genre,
		Rating:   filmSaved.Rating,
		Duration: filmSaved.Duration,
		Quality:  filmSaved.Quality,
		Trailer:  filmSaved.Trailer,
		Watch:    filmSaved.Watch,
	}
}

func (service *FilmServiceImpl) FindById(filmId string) web.FilmResponse {
	db := service.DB
	film, err := service.FilmRepository.FindById(db, filmId)
	if err != nil {
		panic(err)
	}

	return web.FilmResponse{
		Id:       film.Id,
		Title:    film.Title,
		Genre:    film.Genre,
		Rating:   film.Rating,
		Duration: film.Duration,
		Quality:  film.Quality,
		Trailer:  film.Trailer,
		Watch:    film.Watch,
	}
}

func (service *FilmServiceImpl) FindAll(page int) ([]web.FilmResponse, int) {
	db := service.DB
	films, totalPage := service.FilmRepository.FindAll(db, page)
	var filmResponse []web.FilmResponse
	for _, film := range films {
		filmResponse = append(filmResponse, web.FilmResponse{
			Id:       film.Id,
			Title:    film.Title,
			Genre:    film.Genre,
			Rating:   film.Rating,
			Duration: film.Duration,
			Quality:  film.Quality,
			Trailer:  film.Trailer,
			Watch:    film.Watch,
		})
	}
	return filmResponse, totalPage
}

func (service *FilmServiceImpl) Delete(filmId string) {
	db := service.DB
	service.FilmRepository.Delete(db, filmId)
}

func (service *FilmServiceImpl) FindByTitle(page int, title string) ([]web.FilmResponse, int) {
	db := service.DB
	films, totalPage, err := service.FilmRepository.FindByTitle(db, page, title)
	if err != nil {
		panic(err)
	}
	var filmResponse []web.FilmResponse
	for _, film := range films {
		filmResponse = append(filmResponse, web.FilmResponse{
			Id:       film.Id,
			Title:    film.Title,
			Genre:    film.Genre,
			Rating:   film.Rating,
			Duration: film.Duration,
			Quality:  film.Quality,
			Trailer:  film.Trailer,
			Watch:    film.Watch,
		})
	}
	return filmResponse, totalPage
}
