package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	controllers "github.com/wainola/GoReact/controllers"
)

// struct employees => TODO: moverlo a un directorio de structs
type employees struct {
	empNo     int64
	birthDate time.Time
	firstName string
	lastName  string
	gender    string
	hireDate  time.Time
}

// api endpoint para testeos
func GetApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api endpoint")
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "./src")
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/employees", controllers.GetEmployees)
	http.HandleFunc("/api", GetApi)

	log.Println("Listening on port 4500")
	http.ListenAndServe(":4500", nil)
}
