package domain

type EmployeeResponse struct {
	Status       string       `json:"status"`
	EmployeeList EmployeeList `json:"data"`
}

type EmployeeAPI interface {
	EmployeeList() (EmployeeList, error)
}
