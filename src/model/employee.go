package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}
