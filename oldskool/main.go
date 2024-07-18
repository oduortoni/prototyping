package main

import (
	"net/http"

	"github.com/oduortoni/prototyping/stat"
)

func main() {
	println("Serever is running on http://localhost:3500/")

	http.HandleFunc("/", stat.StatIndex)
	http.HandleFunc("/stat", stat.StatServe)
	http.ListenAndServe(":3500", nil)
}
