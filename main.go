package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Start()
	http.HandleFunc("/search", groupie.Search)
	http.HandleFunc("/style/{file}", groupie.Style)
	http.HandleFunc("/", groupie.Home)
	http.HandleFunc("/artist/{id}", groupie.ArtistInfo)
	fmt.Println("http://localhost:8070")
	log.Fatal(http.ListenAndServe(":8070", nil))
}
