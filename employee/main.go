package main

import (
	"context"
	"employee/config"
	"employee/constants"
	"employee/routes"
	"employee/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	server      *gin.Engine
	ctx         context.Context
)

func main() {

	server = gin.Default()
	mongoclient, err := config.ConnectDatabase()
	fmt.Println("hello")
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	service.EmployeeTableCollection = mongoclient.Database(constants.DbName).Collection("Employee_details")
	routes.CustomerTableRoute(server)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
