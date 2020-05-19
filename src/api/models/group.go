package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID           primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Name	     string             	`json:"name"`
	Subjects	 []primitive.ObjectID   `json:"subjects" bson:"subjects"`
	Students	 []primitive.ObjectID 	`json:"students" bson:"students"`
}