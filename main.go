package main

import (
	"log"
	"net/http"

	"github.com/oduortoni/prototyping/stat"
)

func main() {
	http.HandleFunc("/", stat.StatIndex)
	http.HandleFunc("/stat", stat.StatServe)
	http.ListenAndServe(":3000", nil)
	log.Println("Serever is running on http://localhost:3000/")
}
