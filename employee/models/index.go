package models

import "time"

type Employee struct {
	Id          int       `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Dateofbirth string    `json:"dateofbirth" bson:"dateofbirth"`
	CreatedDate time.Time `json:"created_date" bson:"created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date"`
}
type UpdatDetails struct {
	OldName string `bson:"oldname"`
	NewName string `bson:"newname"`
}
type Deleteid struct {
	Id int `json:"id"bson:"id `
}
