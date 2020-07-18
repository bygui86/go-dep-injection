package domain

import (
	"fmt"
)

type Employee struct {
	ID   string `json:"ID"`
	Name string `json:"employee_name"`
	Age  string `json:"employee_age"`
}

type EmployeeList []Employee

func (e *Employee) String() string {
	return fmt.Sprintf("ID: %s - Name: %s - Age: %s", e.ID, e.Name, e.Age)
}
