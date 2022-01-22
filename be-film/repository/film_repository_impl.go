package repository

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/google/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/zenklot/api-film/model/domain"
)

type FilmRepositoryImpl struct {
}

func NewFilmRepository() FilmRepository {
	return &FilmRepositoryImpl{}
}

func (repository *FilmRepositoryImpl) Save(db *scribble.Driver, film domain.Film) domain.Film {
	id := uuid.New().String()
	film.Id = id
	err := db.Write("film", id, film)
	if err != nil {
		panic(err)
	}
	return film
}

func (repository *FilmRepositoryImpl) Delete(db *scribble.Driver, filmId string) {
	err := db.Delete("film", filmId)
	if err != nil {
		panic(err)
	}
}

func (repository *FilmRepositoryImpl) FindAll(db *scribble.Driver, page int) ([]domain.Film, int) {
	records, err := db.ReadAll("film")
	if err != nil {
		panic(err)
	}

	perPage := 15
	totalPage := int(math.Ceil(float64(float64(len(records)) / float64(perPage))))

	awal := (page - 1) * perPage
	akhir := page * perPage
	if len(records)-akhir < 0 {
		akhir = len(records)
	}

	films := []domain.Film{}
	for i := awal; i < akhir; i++ {
		filmData := domain.Film{}
		if err := json.Unmarshal([]byte(records[i]), &filmData); err != nil {
			fmt.Println("Error", err)
		}
		films = append(films, filmData)
	}

	return films, totalPage
}

func (repository *FilmRepositoryImpl) FindById(db *scribble.Driver, filmId string) (domain.Film, error) {
	film := domain.Film{}
	if err := db.Read("film", filmId, &film); err != nil {
		return film, err
	}
	return film, nil
}

func (repository *FilmRepositoryImpl) FindByTitle(db *scribble.Driver, page int, titleFilm string) ([]domain.Film, int, error) {
	records, err := db.ReadAll("film")
	if err != nil {
		panic(err)
	}

	films := []domain.Film{}
	for _, dataFilm := range records {
		filmData := domain.Film{}
		if err := json.Unmarshal([]byte(dataFilm), &filmData); err != nil {
			fmt.Println("Error", err)
		}
		if strings.Contains(strings.ToLower(filmData.Title), strings.ToLower(titleFilm)) {
			films = append(films, filmData)
		}
	}

	perPage := 15
	totalPage := int(math.Ceil(float64(float64(len(films)) / float64(perPage))))
	awal := (page - 1) * perPage
	akhir := page * perPage

	if len(films)-akhir < 0 {
		akhir = len(films)
	}

	filmSearch := []domain.Film{}

	for i := awal; i < akhir; i++ {

		filmSearch = append(filmSearch, films[i])

	}

	return filmSearch, totalPage, err
}
