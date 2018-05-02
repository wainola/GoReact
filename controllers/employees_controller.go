package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wainola/GoReact/types"
)

// variable de configuracion para credenciales de la DB
var configDb map[string]interface{}
var db *sql.DB
var err error

func init() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("dat es %s\n", configDb)
	if err := json.Unmarshal(config, &configDb); err != nil {
		panic(err)
	}
	//fmt.Println(&configDb)

	conn := fmt.Sprintf("%s:%s@/employees", configDb["user"], configDb["password"])

	//fmt.Printf("Cadena de conexion es: %s", conn)

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

// HTTP GET -> todos los empleados limit 100
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	defer db.Close()
	var emp_no int64
	var birth_date []uint8
	var name string
	var lastName string
	var gender []uint8
	var hire_date []uint8
	rows, e := db.Query(
		`select emp_no, birth_date, first_name, last_name, gender, hire_date
		from employees limit 100;`)
	if e != nil {
		log.Println(e)
		return
	}
	var emps []types.Empleados
	for rows.Next() {
		err := rows.Scan(&emp_no, &birth_date, &name, &lastName, &gender, &hire_date)
		if err != nil {
			log.Println(err)
		}
		log.Println(emp_no, string(birth_date[:]), name, lastName, string(gender[:]), string(hire_date[:]))
		layout := "2006-01-02"
		b_date, _ := time.Parse(layout, string(birth_date[:]))
		h_date, _ := time.Parse(layout, string(hire_date[:]))
		e_no := int64(emp_no)
		gen := string(gender[:])
		e := types.Empleados{
			EmpNo:     e_no,
			BirthDate: b_date,
			FirstName: name,
			LastName:  lastName,
			Gender:    gen,
			HireDate:  h_date,
		}
		log.Println(e)
		emps = append(emps, e)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(emps)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP POST -> añadir nuevo empleado
func PostEmployees(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	// Chequeo si el metodo es post para poder hacer la insercion
	if r.Method == "POST" {
		defer db.Close()
		// last_emp_no para insertar el ultimo empleado con su id correcto
		var last_emp_no int64
		fila, e := db.Query(`select emp_no from employees order by emp_no desc limit 1`)
		if e != nil {
			panic(e)
		}
		for fila.Next() {
			err := fila.Scan(&last_emp_no)
			if err != nil {
				log.Fatal("Hay un error!")
				panic(err)
			}
			log.Println("Last emp_no es...")
			log.Println(last_emp_no + 1)
		}
		// leyendo el body del request para poder ingresar el empleado
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body[:]))
	} else {
		d := map[string]string{"msg": "wrong method!"}
		j, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}
