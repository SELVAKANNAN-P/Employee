package models

type Employee struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Date_of_birth string `json:"date___of___birth,omitempty"`
	CreatedDate   string `json:"created_date,omitempty"`
	UpdatedDate   string `json:"updated_date"`
}
type UpdatDetails struct {
	OldName string `bson:"oldname"`
	NewName string `bson:"newname"`
}
type Deleteid struct{
	Id            int    `json:"id,omitempty"`
}