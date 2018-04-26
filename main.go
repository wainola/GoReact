package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type employees struct {
	empNo     int64
	birthDate time.Time
	firstName string
	lastName  string
	gender    string
	hireDate  time.Time
}

// api endpoint test
func GetApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api endpoint")
}

// get all employees limit 100
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var name string
	var lastName string
	rows, e := db.Query(
		`select first_name, last_name
		from employees limit 100;`)
	if e != nil {
		log.Println(e)
		return
	}
	emps := make([]employees, 0)
	for rows.Next() {
		err := rows.Scan(&name, &lastName)
		if err != nil {
			log.Println(err)
		}
		log.Println(name, lastName)
		json.NewEncoder(w).Encode(map[string]string{"nombre": name, "apellido": lastName})
	}
	log.Println(emps)

}

func init() {
	db, err = sql.Open("mysql", "nrriquel:Nrriquel1987@/employees")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	defer db.Close()
	var dir string
	flag.StringVar(&dir, "dir", ".", "./src")
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/employees", GetEmployees)
	http.HandleFunc("/api", GetApi)

	log.Println("Listening on port 4500")
	http.ListenAndServe(":4500", nil)
}
