package config

import (
	"context"
	"employee/constants"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	//"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

func ConnectDatabase() (*mongo.Client, error) {
	ctx := context.TODO()
	mongoconnnection := options.Client().ApplyURI(constants.Connectionstring)
	mongoclient, err := mongo.Connect(ctx, mongoconnnection)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Println("Database Connected")
	return mongoclient, nil

}
