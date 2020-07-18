package interfaces

import (
	"log"

	"github.com/bygui86/go-dingo/dingo-dep-inj/domain"
)

type AppController struct {
	employeeAPI domain.EmployeeAPI
}

func (ac *AppController) New(employeeAPI domain.EmployeeAPI) {
	ac.employeeAPI = employeeAPI
}

// We create a method named 'Inject' to allow Dingo injecting all needed dependencies, defined in 'module.go'
func (ac *AppController) Inject(employeeAPI domain.EmployeeAPI) {
	ac.employeeAPI = employeeAPI
}

func (ac *AppController) PrintEmployees() {
	log.Println("Print employees")
	list, err := ac.employeeAPI.EmployeeList()
	if err != nil {
		log.Fatalf("Error getting employees from %s: %s\n", ac.employeeAPI, err.Error())
	}

	log.Printf("Print %d received employees:\n", len(list))
	for _, l := range list {
		log.Printf("    %s\n", l.String())
	}
}
