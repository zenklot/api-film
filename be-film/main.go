package main

import (
	"io/ioutil"

	"github.com/zenklot/api-film/database/seeder"
)

func main() {
	collection, _ := ioutil.ReadDir("./app/db")
	if len(collection) == 0 {
		seeder.Seed()
	}
}
