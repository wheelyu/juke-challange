package service

import (
	"juke-challange/src/model"
	"juke-challange/src/repository"
)

func GetEmployees() ([]model.Employee, error) {
	return repository.GetAllEmployees()
}

func CreateEmployee(employee *model.Employee) error {
	return repository.CreateEmployee(employee)
}

func UpdateEmployee(id string, employee *model.Employee) error {
	return repository.UpdateEmployee(id, employee)
}

func DeleteEmployee(id string) error {
	return repository.DeleteEmployee(id)
}
