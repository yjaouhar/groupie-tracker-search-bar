package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Isfetched = groupie.Fetch("artists", "")
	http.HandleFunc("/style/{file}", groupie.Style)
	http.HandleFunc("/", groupie.Home)
	http.HandleFunc("/artist/{id}", groupie.ArtistInfo)
	fmt.Println("http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
