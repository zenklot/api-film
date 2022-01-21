package repository

import (
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/zenklot/api-film/model/domain"
)

type FilmRepository interface {
	Save(db *scribble.Driver, film domain.Film) domain.Film
	Delete(db *scribble.Driver, filmId string)
	FindAll(db *scribble.Driver, page int) ([]domain.Film, int)
	FindById(db *scribble.Driver, domainId string) (domain.Film, error)
	FindByTitle(db *scribble.Driver, page int, titleFilm string) ([]domain.Film, int, error)
}
