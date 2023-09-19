package routes

import (
	"employee/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerTableRoute(router *gin.Engine) {
	router.POST("/api/EmployeeTable/create", controllers.CreateEmployeeTable)
	router.GET("/Employee/getbyid", controllers.GetEmployeeById)
	router.PUT("/Employee/ubdatebyid", controllers.UpdateEmployeeTable)
	router.DELETE("/Employee/deleteid", controllers.DeleteEmployeeTable)

}

