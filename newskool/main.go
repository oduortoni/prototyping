package main

import (
	"net/http"

	"github.com/oduortoni/prototyping/jokes"
	"github.com/oduortoni/prototyping/stat"
)

func main() {
	println("Serever is running on http://localhost:4500/")

	http.HandleFunc("/", stat.StatIndex)
	http.HandleFunc("/stat", stat.StatServe)
	http.HandleFunc("/jokes", jokes.Jokes)
	http.ListenAndServe(":4500", nil)
}
