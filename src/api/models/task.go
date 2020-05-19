package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SubjectID	 primitive.ObjectID `json:"subject_id" bson:"subject_id"`
	Name	     string             `json:"name"`
	Description  string             `json:"description"`
	CreatedAt    string				`json:"created_at" bson:"created_at"`
	UpdatedAt    string				`json:"updated_at" bson:"updated_at"`
	Deadline     string				`json:"deadline"`
}
