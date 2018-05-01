package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/wainola/GoReact/controllers"
)

func main() {

	var dir string
	flag.StringVar(&dir, "dir", ".", "./src")
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/employees", controllers.GetEmployees)
	http.HandleFunc("/newEmployee", controllers.PostEmployees)

	log.Println("Listening on port 4500")
	http.ListenAndServe(":4500", nil)
}
