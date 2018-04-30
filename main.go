package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// variable de configuracion para credenciales de la DB
var configDb map[string]interface{}
var db *sql.DB
var err error

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
	var emps []interface{}
	//emps := make([]employees, 0)
	for rows.Next() {
		err := rows.Scan(&name, &lastName)
		if err != nil {
			log.Println(err)
		}
		//log.Println(name, lastName)
		//json.NewEncoder(w).Encode(map[string]string{"nombre": name, "apellido": lastName})
		j, err := json.Marshal(map[string]string{"name": name, "lastName": lastName})
		log.Println(string(j[:]))
		emps = append(emps, j)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(emps)
	//log.Println(emps[0])

}

func init() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("dat es %s\n", configDb)
	if err := json.Unmarshal(config, &configDb); err != nil {
		panic(err)
	}
	fmt.Println(&configDb)

	conn := fmt.Sprintf("%s:%s@/employees", configDb["user"], configDb["password"])

	fmt.Printf("Cadena de conexion es: %s", conn)

	// iniciando conexion a db
	db, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Conexion exitosa a MySQL")
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
