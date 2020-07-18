package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bygui86/go-dingo/dingo-dep-inj/domain"
)

const endpoint = "/api/v1/employees"

type ConcreteEmployeeAPI struct {
	// The inject:"config.productApiUrl" tag will inject the provided string into the attribute.
	ApiUrl string `inject:"config:employeeApi"`
}

func (cpa *ConcreteEmployeeAPI) EmployeeList() (domain.EmployeeList, error) {
	body, getErr := cpa.Get(endpoint)
	if getErr != nil {
		return nil, getErr
	}

	// log.Printf("Unmarshal employee from JSON: %s\n", string(body))
	var resp domain.EmployeeResponse
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return resp.EmployeeList, nil
}

func (cpa *ConcreteEmployeeAPI) Get(endpoint string) ([]byte, error) {
	log.Printf("Get employees from '%s'\n", cpa.ApiUrl)
	resp, err := http.Get(fmt.Sprintf("%s%s", cpa.ApiUrl, endpoint))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
