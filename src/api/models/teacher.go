package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name	     string             `json:"name"`
}