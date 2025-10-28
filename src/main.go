package main

import (
	"juke-challange/src/controller"
	"juke-challange/src/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDB()
	model.DB.AutoMigrate(&model.Employee{})

	r := gin.Default()

	r.GET("/employees", controller.GetEmployees)
	r.POST("/employees", controller.CreateEmployee)
	r.PUT("/employees/:id", controller.UpdateEmployee)
	r.DELETE("/employees/:id", controller.DeleteEmployee)

	r.Run(":8080")
}
