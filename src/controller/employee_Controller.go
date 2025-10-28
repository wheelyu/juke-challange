package controller

import (
	"errors"
	"juke-challange/src/model"
	"juke-challange/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllEmployees godoc
// @Summary Get all employees
// @Description Menampilkan daftar semua karyawan
// @Tags employees
// @Produce json
// @Success 200 {array} model.CreateEmployeeRequest
// @Router /api/employees [get]
func GetEmployees(c *gin.Context) {
	employees, err := service.GetEmployees()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, employees)
}

// GetEmployeeByID godoc
// @Summary Get employee by ID
// @Description Menampilkan detail karyawan berdasarkan ID
// @Tags employees
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} model.CreateEmployeeRequest
// @Router /api/employees/{id} [get]
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := service.GetEmployeeByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Add a new employee to the system
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body model.CreateEmployeeRequest true "Employee Data"
// @Success 200 {object} model.CreateEmployeeRequest
// @Failure 400 {object} map[string]string
// @Router /api/employees [post]
func CreateEmployee(c *gin.Context) {
	var req model.CreateEmployeeRequest

	// Bind JSON ke struct request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// Validasi input
	if req.Name == "" {
		c.Error(errors.New("name cannot be empty"))
		return
	}
	if req.Email == "" {
		c.Error(errors.New("email cannot be empty"))
		return
	}
	if req.Position == "" {
		c.Error(errors.New("position cannot be empty"))
		return
	}
	if req.Salary <= 0 {
		c.Error(errors.New("salary must be greater than zero"))
		return
	}

	// Buat objek employee dari request
	employee := model.Employee{
		Name:     req.Name,
		Email:    req.Email,
		Position: req.Position,
		Salary:   req.Salary,
	}

	// Simpan ke database via service
	if err := service.CreateEmployee(&employee); err != nil {
		c.Error(err)
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee created successfully",
		"id":      employee.ID,
	})
}

// UpdateEmployee godoc
// @Summary Update an existing employee
// @Description Update the details of an existing employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body model.CreateEmployeeRequest true "Employee Data"
// @Success 200 {object} model.CreateEmployeeRequest
// @Failure 400 {object} map[string]string
// @Router /api/employees/{id} [put]
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

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Menghapus karyawan berdasarkan ID
// @Tags employees
// @Param id path string true "Employee ID"
// @Success 200 {object} map[string]string
// @Router /api/employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteEmployee(id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
