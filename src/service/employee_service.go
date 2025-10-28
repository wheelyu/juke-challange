package service

import (
	"juke-challange/src/model"
	"juke-challange/src/repository"

	"errors"
)

func GetEmployees() ([]model.Employee, error) {
	return repository.GetAllEmployees()
}

func GetEmployeeByID(id string) (*model.Employee, error) {
	return repository.GetEmployeeByID(id)
}

func CreateEmployee(employee *model.Employee) error {

	result := model.DB.Where("email = ?", employee.Email).First(&employee)
	if result.RowsAffected > 0 {
		return errors.New("email already exists")
	}

	return repository.CreateEmployee(employee)
}

func UpdateEmployee(id string, employee *model.Employee) error {
	return repository.UpdateEmployee(id, employee)
}

func DeleteEmployee(id string) error {
	return repository.DeleteEmployee(id)
}
