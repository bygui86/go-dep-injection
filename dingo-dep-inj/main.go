package main

import (
	"log"

	"flamingo.me/dingo"

	"github.com/bygui86/go-dingo/dingo-dep-inj/domain"
	"github.com/bygui86/go-dingo/dingo-dep-inj/infrastructure"
	"github.com/bygui86/go-dingo/dingo-dep-inj/interfaces"
)

const (
	employeeApiDingoConfig = "config:employeeApi"
	employeeApi            = "http://dummy.restapiexample.com"
)

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	// For basic types like string or integer values we can just use AnnotatedWith and pass a key
	injector.Bind(new(string)).AnnotatedWith(employeeApiDingoConfig).ToInstance(employeeApi)

	// We bind infrastructure.ConcreteEmployeeAPI to all domain.EmployeeAPI
	injector.Bind(new(domain.EmployeeAPI)).To(infrastructure.ConcreteEmployeeAPI{})
}

func main() {
	log.Println("Start 'dingo-dep-inj' sample")

	log.Println("Create Dingo injector")
	injector, dingoErr := dingo.NewInjector(new(Module))
	if dingoErr != nil {
		log.Fatalf("Error creating Dingo injector: %s\n", dingoErr)
	}

	// Now we can build objects with the Dingo injector:
	// 1. get a new instance of AppController
	log.Println("Get a new instance of AppController")
	appCtrlInstance, appCtrlErr := injector.GetInstance(new(interfaces.AppController))
	if appCtrlErr != nil {
		log.Fatalf("Error getting AppController from Dingo injector: %s\n", appCtrlErr)
	}

	// 2. cast instance accordingly to AppController
	log.Println("Cast instance accordingly to AppController")
	appController := appCtrlInstance.(*interfaces.AppController)

	// 3. use AppController appCtrlInstance normally
	appController.PrintEmployees()

	log.Println("'dingo-dep-inj' sample completed")
}
