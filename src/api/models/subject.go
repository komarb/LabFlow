package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subject struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Teacher      Teacher            `json:"teacher"`
	Name  string             		`json:"name"`
	Groups       []string    		`json:"groups"`
}
