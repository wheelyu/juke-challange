package repository

import (
	"juke-challange/src/model"
)

func GetAllEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	result := model.DB.Find(&employees)
	return employees, result.Error
}

func CreateEmployee(employee *model.Employee) error {
	return model.DB.Create(employee).Error
}

func UpdateEmployee(id string, updated *model.Employee) error {
	var emp model.Employee
	if err := model.DB.First(&emp, id).Error; err != nil {
		return err
	}
	emp.Name = updated.Name
	emp.Email = updated.Email
	emp.Position = updated.Position
	emp.Salary = updated.Salary
	return model.DB.Save(&emp).Error
}

func DeleteEmployee(id string) error {
	return model.DB.Delete(&model.Employee{}, id).Error
}
