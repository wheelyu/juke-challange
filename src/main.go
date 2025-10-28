package main

import (
	"juke-challange/src/middleware"
	"juke-challange/src/model"
	"juke-challange/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	model.ConnectDB()
	model.DB.AutoMigrate(&model.Employee{})

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	routes.SetupRoutes(r)

	r.Run(":8080")
}
