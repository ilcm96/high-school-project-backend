package model

type Todo struct {
	Id   string `json:"_id" bson:"_id"`
	Todo string `json:"todo"`
}
