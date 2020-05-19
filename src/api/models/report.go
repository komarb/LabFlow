package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TaskID		 primitive.ObjectID `json:"taskId" bson:"taskId"`
	SubjectID	 primitive.ObjectID `json:"subjectId" bson:"subjectId"`
	ReporterID   primitive.ObjectID `json:"reporterId" bson:"reporterId"`
	Date         string             `json:"date"`
	Text         string             `json:"text"`
	TeachersNote string             `json:"teachers_note" bson:"teachers_note"`
	State        string 		    `json:"state"`
	Archived     bool               `json:"archived"`
	Reporter	 []User				`json:"reporter"`
}
