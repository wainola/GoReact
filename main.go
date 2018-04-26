package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":4500", nil)
}
