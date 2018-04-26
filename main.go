package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func GetApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api endpoint")
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "./src")
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/api", GetApi)

	log.Println("Listening on port 4500")
	http.ListenAndServe(":4500", nil)
}
