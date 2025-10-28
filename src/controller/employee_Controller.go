package controller

import (
	"errors"
	"juke-challange/src/model"
	"juke-challange/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	employees, err := service.GetEmployees()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := service.GetEmployeeByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func CreateEmployee(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.Error(err)
		return
	}
	if employee.Name == "" {
		c.Error(errors.New("name cannot be empty"))
		return
	}
	if employee.Email == "" {
		c.Error(errors.New("email cannot be empty"))
		return
	}
	if employee.Position == "" {
		c.Error(errors.New("position cannot be empty"))
		return
	}
	if employee.Salary <= 0 {
		c.Error(errors.New("salary cannot be negative"))
		return
	}
	if err := service.CreateEmployee(&employee); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee created", "id": employee.ID})
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.Error(err)
		return
	}
	if err := service.UpdateEmployee(id, &employee); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated", "id": id})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteEmployee(id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
