package main

import (
	"fmt"
	"log"
	"net/http"
)

const omdbUrl = "http://www.omdbapi.com"

func SearchMovie(title string) {
	// search movie by title
	resp, err := http.Get(fmt.Sprintf("%s/?t=%s", omdbUrl, title))
	if err != nil {
		log.Default().Println(err)
	}
	defer resp.Body.Close()

}
