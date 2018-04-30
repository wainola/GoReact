package types

import "time"

type Empleados struct {
	EmpNo     int64     `json:"EmpNo"`
	BirthDate time.Time `json:"BirthDate"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Gender    rune      `json:"Gender"`
	HireDate  time.Time `json:"HireDate"`
}
