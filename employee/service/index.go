package service

import (
	"context"
	"employee/models"
	"time"

	//
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	EmployeeTableCollection *mongo.Collection
	ctx                     context.Context
)

// func CreateEmployeeTable(user *models.Employee) (string, error) {
//     ctx := context.TODO()

//     _, err := EmployeeTableCollection.InsertOne(ctx, user)
//     if err != nil {
//         return "error", err
//     }
//     return "success", nil
// }


func DeleteEmployeeTable(id *models.Deleteid) (*mongo.DeleteResult, error) {
	d := id.Id
	filter := bson.M{"id": d}
	result, err := EmployeeTableCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func GetEmployeeById(id *models.Deleteid) (*models.Employee, error) {
	//	filter := bson.D{{Key: "id", Value: id}}
	var Employee *models.Employee
	d := id.Id
	filter := bson.M{"id": d}
	res := EmployeeTableCollection.FindOne(ctx, filter)
	err := res.Decode(&Employee)
	if err != nil {
		return nil, err
	}
	return Employee, nil
}

func GetEmployeeByI(id int) (*models.Employee, error) {
	//	filter := bson.D{{Key: "id", Value: id}}
	var Employee *models.Employee
	d := id
	filter := bson.M{"id": d}
	res := EmployeeTableCollection.FindOne(ctx, filter)
	err := res.Decode(&Employee)
	if err != nil {
		return nil, err
	}
	return Employee, nil
}
func CreateEmployeeTable(user *models.Employee) (string, error) {
	ctx := context.TODO()

	if user.Id <=0 {
		fmt.Println("1")
        return "fillallfields", nil
    }
    
    if user.Name == "" {
		fmt.Println("2")
        return "fillallfields", nil
    }
    
    if user.Dateofbirth == "" {
		fmt.Println("4")
        return "fillallfields", nil
    }

	a, err := GetEmployeeByI(user.Id)
	// if err != nil {
	// 	return "error", err
	// }

	if a != nil && a.Id == user.Id {
		return "usedid", nil
	}
	
      user.CreatedDate=time.Now()
	_, err = EmployeeTableCollection.InsertOne(ctx, user)
	if err != nil {
		return "error", err
	}

	return "success", nil
}
// func UpdateemployeerTable(oldname string, newname string) (*mongo.UpdateResult, error) {
// 	filter := bson.D{{Key: "name", Value: oldname}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newname}}}}
// 	result, err := EmployeeTableCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	return result, nil
// }
func UpdateemployeerTable(oldname string, newname string) (*mongo.UpdateResult, error) {
    ctx := context.TODO()

    // Set the UpdatedDate field to the current time
    UpdatedDate := time.Now()

    filter := bson.D{{Key: "name", Value: oldname}}
    update := bson.D{
        {Key: "$set", Value: bson.D{{Key: "name", Value: newname}}},
        {Key: "$set", Value: bson.D{{Key: "updated_date", Value: UpdatedDate}}},
    }

    result, err := EmployeeTableCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    return result, nil
}
// func UpdateEmployee(id int, newName string, newDateofbirth string) (*mongo.UpdateResult, error) {
//     ctx := context.TODO()

//     // Set the UpdatedDate field to the current time
//     updatedTime := time.Now()

//     filter := bson.D{{Key: "id", Value: id}}
//     update := bson.D{
//         {Key: "$set", Value: bson.D{{Key: "name", Value: newName}}},
//         {Key: "$set", Value: bson.D{{Key: "datesofbirth", Value: newDateofbirth}}},
//         {Key: "$set", Value: bson.D{{Key: "updated_date", Value: updatedTime}}},
//     }

//     result, err := EmployeeTableCollection.UpdateOne(ctx, filter, update)
//     if err != nil {
//         fmt.Println(err.Error())
//         return nil, err
//     }
//     return result, nil
// }

