package service

import (
	"context"
	"employee/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	EmployeeTableCollection *mongo.Collection
	ctx                     context.Context
)

func CreateCustomerTable(user *models.Employee) (string, error) {
	ctx = context.TODO()
	_, err := EmployeeTableCollection.InsertOne(ctx, user)
	if err != nil {
		return "error", err
	}
	return "success", nil
}
func UpdateemployeerTable(oldname string, newname string) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "name", Value: oldname}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newname}}}}
	result, err := EmployeeTableCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil
}
func DeleteEmployeeTable(id *models.Deleteid) (*mongo.DeleteResult, error) {
    cust_id := id.Id
    filter := bson.D{{Key: "id", Value: cust_id}}
    result, err := EmployeeTableCollection.DeleteOne(ctx, filter)
    if err != nil {
        return nil, err
    }
    return result, nil
}
