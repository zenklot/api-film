package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/zenklot/api-film/app"
	"github.com/zenklot/api-film/controller"
	"github.com/zenklot/api-film/database/seeder"
	"github.com/zenklot/api-film/repository"
	"github.com/zenklot/api-film/service"
)

func main() {
	collection, _ := ioutil.ReadDir("./app/db")
	if len(collection) == 0 {
		seeder.Seed()
	}

	// connection with scribble to db directory
	db, err := scribble.New("./app/db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	validate := validator.New()
	filmRepository := repository.NewFilmRepository()
	filmService := service.NewFilmService(filmRepository, db, validate)
	filmController := controller.NewFilmController(filmService)

	router := app.NewRouter(filmController)

	address := "localhost:3000"
	server := http.Server{
		Addr:    address,
		Handler: router,
	}
	fmt.Printf("Server Started at %s\n", address)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
