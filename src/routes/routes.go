package routes

import (
	"juke-challange/src/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("api/employees", controller.GetEmployees)
	r.GET("api/employees/:id", controller.GetEmployeeByID)
	r.POST("api/employees", controller.CreateEmployee)
	r.PUT("api/employees/:id", controller.UpdateEmployee)
	r.DELETE("api/employees/:id", controller.DeleteEmployee)
}
