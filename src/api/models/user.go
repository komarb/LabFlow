package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           	primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username		string		`json:"username"`
	Password		string `json:"password"`
	Name			string	`json:"name"`
	Role   			string	`json:"role"`
	Groups			[]string	`json:"groups"`
}

type UserCredentials struct {
	Username		string		`json:"username"`
	Password		string `json:"password"`
}
type AuthResponse struct {
	ID				string		`json:"id"`
	AccessToken		string		`json:"accessToken"`
	Username		string		`json:"username"`
	Name 			string		`json:"name"`
	Role			string		`json:"role"`
	Groups			[]string	`json:"groups"`
}