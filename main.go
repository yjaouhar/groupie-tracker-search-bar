package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	// var mo interface{}
	groupie.Isfetched = groupie.Fetch("artists")
	groupie.Isfetched = groupie.Fetch("relation")
	groupie.Uni()
	// fmt.Println(groupie.Result.Aloo)
	http.HandleFunc("/search", groupie.Search)
	http.HandleFunc("/style/{file}", groupie.Style)
	http.HandleFunc("/", groupie.Home)
	http.HandleFunc("/artist/{id}", groupie.ArtistInfo)
	// http.Handle("hh", mo)
	fmt.Println("http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
