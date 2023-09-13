package controllers

import (
	"employee/models"
	"employee/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployeeTable(ctx *gin.Context) {

	var profile *models.Employee
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := service.CreateCustomerTable(profile)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}
func UpdateAccountsTable(ctx *gin.Context) {
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
	result, err := service.DeleteEmployeeTable(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}
