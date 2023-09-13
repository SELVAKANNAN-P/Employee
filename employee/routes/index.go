package routes

import (
	"employee/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerTableRoute(router *gin.Engine) {
	router.POST("/api/EmployeeTable/create", controllers.CreateEmployeeTable)

}
