package main

import (
	"log"

	"github.com/bygui86/go-dingo/no-dep-inj/infrastructure"
	"github.com/bygui86/go-dingo/no-dep-inj/interfaces"
)

const employeeApi = "http://dummy.restapiexample.com"

func main() {
	log.Println("Start 'no-dep-inj' sample")

	log.Printf("Create 'ConcreteEmployeeAPI' with URL '%s'\n", employeeApi)
	productAPI := infrastructure.ConcreteEmployeeAPI{
		ApiUrl: employeeApi,
	}

	log.Println("Create 'AppController'")
	appController := interfaces.AppController{}
	appController.New(&productAPI)

	appController.PrintEmployees()

	log.Println("'no-dep-inj' sample completed")
}
