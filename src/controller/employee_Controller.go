package controller

import (
	"juke-challange/src/model"
	"juke-challange/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	employees, err := service.GetEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UpdateEmployee(id, &employee); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
