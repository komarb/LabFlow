package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           	primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Login			string		`json:"login"`
	Password		string `json:"password"`
	Name			string	`json:"name"`
	Role   			string	`json:"role"`
	Groups			[]string	`json:"groups"`
}