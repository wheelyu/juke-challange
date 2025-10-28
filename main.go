package main

import (
	_ "juke-challange/docs"
	"juke-challange/src/middleware"
	"juke-challange/src/model"
	"juke-challange/src/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	model.ConnectDB()
	model.DB.AutoMigrate(&model.Employee{})

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	routes.SetupRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
