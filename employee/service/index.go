package service

import (
	"context"
	"employee/models"
	"time"

	//
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	EmployeeTableCollection *mongo.Collection
	ctx                     context.Context
)

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
func UpdateemployeerTable(oldname string, newname string) (*mongo.UpdateResult, error) {
	ctx := context.TODO()

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

func CreateEmployeeTable(user *models.Employee) (string, error) {
	ctx := context.TODO()

	if user.Id <= 0 {
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

     user.CreatedDate = time.Now()
	_, err = EmployeeTableCollection.InsertOne(ctx, user)
	if err != nil {
		return "error", err
	}

	return "success", nil
}































































// func CreateEmployeeTable(user *models.Employee) (string, error) {
//     ctx := context.TODO()

//     _, err := EmployeeTableCollection.InsertOne(ctx, user)
//     if err != nil {
//         return "error", err
//     }
//     return "success", nil
// }

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

//	    result, err := EmployeeTableCollection.UpdateOne(ctx, filter, update)
//	    if err != nil {
//	        fmt.Println(err.Error())
//	        return nil, err
//	    }
//	    return result, nil
//	}f
//	func UpdateemployeerTable(oldname string, newname string) (*mongo.UpdateResult, error) {
//		filter := bson.D{{Key: "name", Value: oldname}}
//		update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newname}}}}
//		result, err := EmployeeTableCollection.UpdateOne(ctx, filter, update)
//		if err != nil {
//			fmt.Println(err.Error())
//			return nil, err
//		}
//		return result, nil
//	}

// // GetAllEmployees retrieves all employee details from the database.
// func GetAllEmployees() ([]models.Employee, error) {
// 	ctx := context.TODO()

// 	// Define an empty filter to match all documents
// 	filter := bson.M{}

// 	// Create options for sorting or filtering, if needed
// 	options := options.Find()

// 	cursor, err := EmployeeTableCollection.Find(ctx, filter, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var employees []models.Employee

// 	for cursor.Next(ctx) {
// 		var employeeDoc struct {
// 			Id          int     `bson:"id"`
// 			Name        string  `bson:"name"`
// 			Dateofbirth string  `bson:"dateofbirth"`
// 			CreatedDate *string `bson:"created_date"`
// 			UpdatedDate *string `bson:"updated_date"`
// 		}

// 		if err := cursor.Decode(&employeeDoc); err != nil {
// 			return nil, err
// 		}

// 		employee := models.Employee{
// 			Id:          employeeDoc.Id,
// 			Name:        employeeDoc.Name,
// 			Dateofbirth: employeeDoc.Dateofbirth,
// 		}

// 		if employeeDoc.CreatedDate != nil {
// 			createdDate, err := time.Parse(time.RFC3339, *employeeDoc.CreatedDate)
// 			if err != nil {
// 				return nil, err
// 			}
// 			employee.CreatedDate = &createdDate
// 		}

// 		if employeeDoc.UpdatedDate != nil {
// 			updatedDate, err := time.Parse(time.RFC3339, *employeeDoc.UpdatedDate)
// 			if err != nil {
// 				return nil, err
// 			}
// 			employee.UpdatedDate = &updatedDate
// 		}

// 		employees = append(employees, employee)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return employees, nil
// }

