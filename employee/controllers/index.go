package controllers

import (
	"employee/models"
	"employee/service"
	"net/http"


	// "strconv"

	"github.com/gin-gonic/gin"
)

// func CreateEmployeeTable(ctx *gin.Context) {
//     var profile models.Employee
//     if err := ctx.ShouldBindJSON(&profile); err != nil {
//         ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
//         return
//     }

//     status, err := service.CreateEmployeeTable(&profile)
//     if err != nil {
//         ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
//         return
//     }

//     ctx.JSON(http.StatusCreated, gin.H{"status": status})
// }

func UpdateEmployeeTable(ctx *gin.Context) {
	var profile *models.UpdatDetails
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	customer, err := service.UpdateemployeerTable(profile.OldName, profile.NewName)
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": customer})

}
func DeleteEmployeeTable(ctx *gin.Context) {
	var id *models.Deleteid
	if err := ctx.ShouldBindJSON(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result, err := service.DeleteEmployeeTable(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}

func GetEmployeeById(ctx *gin.Context) {
	var id *models.Deleteid
	if err := ctx.ShouldBindJSON(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	val, err := service.GetEmployeeById(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

// func CreateEmployeeTable(ctx *gin.Context) {
// 	var profile models.Employee
// 	if err := ctx.ShouldBindJSON(&profile); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
// 		return
// 	}

// 	newProfile, err := service.CreateEmployee(&profile)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Employee with the same name already exists"})
// 		} else {
// 			ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 		}
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
// }

// func GetEmployeeByNameHandler(ctx *gin.Context) {
//     name := ctx.Param("name")
//     employee, err := service.GetEmployeeByName(name)
//     if err != nil {
//         ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
//         return
//     }

//     ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": employee})
// }

// Add more controller functions for retrieving, updating, and deleting employees as needed.
func CreateEmployeeTable(ctx *gin.Context) {
	var profile models.Employee
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	status, err := service.CreateEmployeeTable(&profile)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": status})
}
