package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

// get all employees limit 100
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
	// type Empleado struct {
	// 	Name     string `json:"name"`
	// 	LastName string `json:"lastName"`
	// }
	//var emps []types.Empleados
	//emps := make([]employees, 0)
	for rows.Next() {
		err := rows.Scan(&emp_no, &birth_date, &name, &lastName, &gender, &hire_date)
		if err != nil {
			log.Println(err)
		}
		log.Println(emp_no, string(birth_date[:]), name, lastName, string(gender[:]), string(hire_date[:]))
		layout := "01-01-2006"
		date1, err := time.Parse(layout, string(birth_date[:]))
		if err != nil {
			panic(err)
		}
		log.Println(date1)
		//json.NewEncoder(w).Encode(map[string]string{"nombre": name, "apellido": lastName})
		// e := types.Empleados{
		// 	EmpNo:     emp_no,
		// 	BirthDate: birth_date,
		// 	FirstName: name,
		// 	LastName:  lastName,
		// 	Gender:    gender,
		// 	HireDate:  hire_date,
		// }
		//log.Println(e)
		//j, err := json.Marshal(map[string]string{"name": name, "lastName": lastName})
		//log.Println(string(j[:]))
		//emps = append(emps, e)
	}
	// w.Header().Set("Content-Type", "application/json")
	// j, err := json.Marshal(emps)
	// if err != nil {
	// 	panic(err)
	// }
	// w.WriteHeader(http.StatusOK)
	// log.Println(emps)
	// log.Println(j)
	// w.Write(j)
	//log.Println(emps[0])

}
