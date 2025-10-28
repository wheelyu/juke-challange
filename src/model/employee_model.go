package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string
	Email    string
	Position string
	Salary   float64
}

type CreateEmployeeRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Position string  `json:"position" binding:"required"`
	Salary   float64 `json:"salary" binding:"required,gte=0"`
}
