package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/zenklot/api-film/model/domain"
)

type Data struct {
	Data []Film `json:"data"`
}

type Film struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Genre    []string `json:"genre"`
	Rating   string   `json:"rating"`
	Duration string   `json:"duration"`
	Quality  string   `json:"quality"`
	Trailer  string   `json:"trailer"`
	Watch    string   `json:"watch"`
}

func Seed() {
	db, err := scribble.New("./app/db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	// Open json file
	jsonFile, err := os.Open("database/seeder/dummy.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Seeding Started from dummy.json")

	defer jsonFile.Close()

	// Convert json file to []byte for process unmarshal json to struct
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Data
	json.Unmarshal(byteValue, &data)

	// save to db using for
	for _, film := range data.Data {
		id := uuid.New().String()
		genre := film.Genre
		for i := range genre {
			genre[i] = strings.TrimSpace(genre[i])
		}

		db.Write("film", id, domain.Film{
			Id:       id,
			Title:    strings.TrimSpace(film.Title),
			Genre:    genre,
			Rating:   strings.TrimSpace(film.Rating),
			Duration: strings.TrimSpace(film.Duration),
			Quality:  film.Quality,
			Trailer:  film.Trailer,
			Watch:    film.Watch})
	}

	fmt.Println("Seeding success from dummy.json")

}
