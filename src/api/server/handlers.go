package server

import (
	"LabFlow/models"
	"LabFlow/utils"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

/*func getAllReports(w http.ResponseWriter, r *http.Request) {
	reports := make([]models.Report, 0)
	var filter bson.M

	

	switch {
	case isAdmin():
		filter = bson.M{"archived" : false}
	case isUser():
		filter = bson.M{"reportsender": Claims.Sub, "archived" : false}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Find",
			"handler" : "getAllReports",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &reports)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getAllReports",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	json.NewEncoder(w).Encode(reports)
}

func getAllReportsSorted(w http.ResponseWriter, r *http.Request) {
	var filter bson.M
	reports := make([]models.Report, 0)

	

	switch {
	case isAdmin():
		filter = bson.M{"archived" : false}
	case isUser():
		filter = bson.M{"reportsender": Claims.Sub, "archived" : false}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	data := mux.Vars(r)
	sortVar := data["var"]
	findOptions := options.Find()
	switch sortVar {
	case "name":
		findOptions.SetSort(bson.M{"reportsender": 1})
	case "date":
		findOptions.SetSort(bson.M{"date": 1})
	}

	cur, err := userCollection.Find(ctx, filter, findOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Find",
			"handler" : "getAllReportsSorted",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &reports)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getAllReports",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	json.NewEncoder(w).Encode(reports)
}

func getReport(w http.ResponseWriter, r *http.Request) {
	var filter bson.M
	var report models.Report

	

	json.NewDecoder(r.Body).Decode(&report)
	data := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(string(data["id"]))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	filter = bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = userCollection.FindOne(ctx, filter).Decode(&report)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if report.ReportSender == Claims.Sub || isAdmin() {
		json.NewEncoder(w).Encode(report)
	} else {
		w.WriteHeader(403)
		return
	}
}

func getArchivedReports(w http.ResponseWriter, r *http.Request) {
	reports := make([]models.Report, 0)
	var filter bson.M

	

	switch {
	case isAdmin():
		filter = bson.M{"archived" : true}
	case isUser():
		filter = bson.M{"reportsender": Claims.Sub, "archived" : true}
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Find",
			"handler" : "getArchievedReports",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &reports)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getArchievedReports",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	json.NewEncoder(w).Encode(reports)
}

func getEmployeeReports(w http.ResponseWriter, r *http.Request) {
	var filter bson.D
	reports := make([]models.Report, 0)

	

	data := mux.Vars(r)
	employee := data["employee"]
	if employee != Claims.Sub && !isAdmin() {
		w.WriteHeader(403)
		return
	}

	if data["dateBegin"] != "" && data["dateEnd"] != "" {
		dateBegin := utils.FormatQueryDate(data["dateBegin"])+"T00:00:00"
		dateEnd := utils.FormatQueryDate(data["dateEnd"])+"T23:59:59"
		filter = bson.D{
			{"reportsender" ,employee},
			{"archived" , false},
			{"$and", []interface{}{
				bson.D{{"date",bson.M{"$gte": dateBegin}}},
				bson.D{{"date", bson.M{"$lte" : dateEnd}}},
			}},
		}
	} else {
		filter = bson.D{
			{"reportsender" ,employee},
			{"archived" , false},
		}
	}

	findOptions := options.Find().SetSort(bson.M{"date": 1})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := userCollection.Find(ctx, filter, findOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Find",
			"handler" : "getEmployeeSampleDate",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &reports)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getEmployeeSampleDate",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	json.NewEncoder(w).Encode(reports)
}


func createReport(w http.ResponseWriter, r *http.Request) {
	var report models.Report

	

	json.NewDecoder(r.Body).Decode(&report)
	report.ReportSender = Claims.Sub
	report.Date = time.Now().Format("2006-01-02T15:04:05")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := userCollection.InsertOne(ctx, report)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.InsertOne",
			"handler" : "createReport",
			"error"	:	err,
		},
		).Fatal("DB interaction resulted in error, shutting down...")
	}
	id := result.InsertedID
	report.ID, err = primitive.ObjectIDFromHex(id.(primitive.ObjectID).Hex())
	json.NewEncoder(w).Encode(report)
}

func updateReport(w http.ResponseWriter, r *http.Request) {
	var report models.Report
	var updatedReport models.Report

	

	json.NewDecoder(r.Body).Decode(&report)
	data := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(string(data["id"]))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	filter := bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = userCollection.FindOne(ctx, filter).Decode(&updatedReport)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if updatedReport.ReportSender == Claims.Sub || isAdmin() {
		updatedReport.Text = report.Text
		if report.Text == "" {
			updatedReport.Archived = true
		} else {
			updatedReport.Archived = false
		}
		updateResult, err := userCollection.ReplaceOne(ctx, filter, updatedReport)
		if err != nil || updateResult.MatchedCount == 0 {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(updatedReport)
	} else {
		w.WriteHeader(403)
		return
	}
}

func deleteReport(w http.ResponseWriter, r *http.Request) {
	var report models.Report
	data := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(string(data["id"]))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	filter := bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = userCollection.FindOne(ctx, filter).Decode(&report)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if report.ReportSender == Claims.Sub || isAdmin() {
		report.Archived = true
		updateResult, err := userCollection.ReplaceOne(ctx, filter, report)
		if err != nil || updateResult.MatchedCount == 0 {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(200)

	} else {
		w.WriteHeader(403)
		return
	}
}
*/

func getSubjects(w http.ResponseWriter, r *http.Request) {
	subjects := make([]models.Subject,0)
	var filter bson.M

	objID, err := primitive.ObjectIDFromHex(Claims.Sub)
	switch {
	case isTeacher():
		filter = bson.M{"teacher._id" : objID}
	case isUser():
		filter = bson.M{"groups" : bson.M{"$in" : Claims.Groups}}
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"name": 1})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := subjectsCollection.Find(ctx, filter, findOptions)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &subjects)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getSubjects",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(subjects)
}
func getSubject(w http.ResponseWriter, r *http.Request) {
	var filter bson.M
	var subject models.Subject

	

	data := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(data["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	filter = bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = subjectsCollection.FindOne(ctx, filter).Decode(&subject)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch {
	case isTeacher():
		objID, err = primitive.ObjectIDFromHex(Claims.Sub)
		if subject.Teacher.ID != objID {
			w.WriteHeader(403)
			return
		}
	case isUser():
		if len(utils.Intersect(Claims.Groups, subject.Groups)) == 0 {
			w.WriteHeader(403)
			return
		}
	}
	json.NewEncoder(w).Encode(subject)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if !isTeacher() {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&task)
	task.CreatedAt = time.Now().Format("2006-01-02T15:04:05")
	task.SubjectID, _ = primitive.ObjectIDFromHex(data["subjectID"])
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := tasksCollection.InsertOne(ctx, task)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.InsertOne",
			"handler" : "createTask",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := result.InsertedID
	task.ID, err = primitive.ObjectIDFromHex(id.(primitive.ObjectID).Hex())
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if !isTeacher() {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data := mux.Vars(r)
	objTaskID, _ := primitive.ObjectIDFromHex(data["taskID"])
	filter := bson.M{"_id": objTaskID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := tasksCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.DeleteOne",
			"handler" : "deleteTask",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getSubjectTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]models.Task, 0)
	var filter bson.M

	

	data := mux.Vars(r)
	objSubjectID, _ := primitive.ObjectIDFromHex(data["subjectID"])
	filter = bson.M{"subject_id" : objSubjectID}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := tasksCollection.Find(ctx, filter)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &tasks)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getSubjectTasks",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func createReport(w http.ResponseWriter, r *http.Request) {
	var report models.Report
	var task models.Task
	
	

	data := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&report)
	objSub, _ := primitive.ObjectIDFromHex(Claims.Sub)
	report.ReporterID = objSub
	report.Date = time.Now().Format("2006-01-02T15:04:05")

	objTaskID, _ := primitive.ObjectIDFromHex(data["taskID"])
	filter := bson.M{"_id": objTaskID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	finderr := tasksCollection.FindOne(ctx, filter).Decode(&task)
	if finderr != nil {
		http.NotFound(w, r)
		return
	}
	report.TaskID = objTaskID

	objSubjectID, _ := primitive.ObjectIDFromHex(data["subjectID"])
	filter = bson.M{"subject_id": objSubjectID}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	finderr = tasksCollection.FindOne(ctx, filter).Decode(&task)
	if finderr != nil {
		http.NotFound(w, r)
		return
	}
	report.SubjectID = objSubjectID

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, err := reportsCollection.InsertOne(ctx, report)
	if err != nil {
		log.WithFields(log.Fields{
			"function": "mongo.InsertOne",
			"handler":  "createReport",
			"error":    err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := result.InsertedID
	report.ID, err = primitive.ObjectIDFromHex(id.(primitive.ObjectID).Hex())
	json.NewEncoder(w).Encode(report)
}

func getReports(w http.ResponseWriter, r *http.Request) {
	reports := make([]models.Report, 0)

	
	data := mux.Vars(r)
	objStudentID, _ := primitive.ObjectIDFromHex(data["studentID"])
	pipeline := []interface{}{
		bson.M{"$lookup" : bson.M{
			"from": "users",
			"localField": "reporterId",
			"foreignField": "_id",
			"as": "reporter",
		}},
		bson.M{"$match":bson.M{
			"reporterId" : objStudentID,
		}},
		bson.M{"$sort":bson.M{
			"date" : -1,
		}},
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := reportsCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Aggregate",
			"handler" : "getReports",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &reports)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getReports",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reports)
}

func getTaskReports(w http.ResponseWriter, r *http.Request) {
	var pipeline []interface{}
	reports := make([]models.Report, 0)

	data := mux.Vars(r)
	log.Info(Claims.Sub)
	objStudentID, _ := primitive.ObjectIDFromHex(Claims.Sub)
	objSubjectID, _ := primitive.ObjectIDFromHex(data["subjectID"])
	objTaskID, _ := primitive.ObjectIDFromHex(data["taskID"])
	if isTeacher() {
		pipeline = []interface{}{
			bson.M{"$lookup" : bson.M{
				"from": "users",
				"localField": "reporterId",
				"foreignField": "_id",
				"as": "reporter",
			}},
			bson.M{"$match":bson.M{
				"subjectId" : objSubjectID,
				"taskId" : objTaskID,
			}},
			bson.M{"$sort":bson.M{
				"date" : -1,
			}},
		}
	} else if isUser() {
		pipeline = []interface{}{
			bson.M{"$lookup" : bson.M{
				"from": "users",
				"localField": "reporterId",
				"foreignField": "_id",
				"as": "reporter",
			}},
			bson.M{"$match":bson.M{
				"subjectId" : objSubjectID,
				"taskId" : objTaskID,
				"reporterId" : objStudentID,
			}},
			bson.M{"$sort":bson.M{
				"date" : -1,
			}},
		}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := reportsCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Aggregate",
			"handler" : "getTaskReports",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	defer cur.Close(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &reports)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getTaskReports",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error, shutting down...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reports)
}

func getSubjectGroups(w http.ResponseWriter, r *http.Request) {
	groups := make([]models.Group, 0)
	var filter bson.M

	

	data := mux.Vars(r)
	objSubjectID, _ := primitive.ObjectIDFromHex(data["subjectID"])
	filter = bson.M{"subjects" : objSubjectID}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := groupsCollection.Find(ctx, filter)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = cur.All(ctx, &groups)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.All",
			"handler" : "getSubjectTasks",
			"error"	:	err,
		},
		).Warn("DB interaction resulted in error!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(groups)
}

func updateReport(w http.ResponseWriter, r *http.Request) {
	var updatedReport models.Report

	data := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(data["reportID"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	filter := bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = reportsCollection.FindOne(ctx, filter).Decode(&updatedReport)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if isTeacher() {
		json.NewDecoder(r.Body).Decode(&updatedReport)
		if updatedReport.Text == "" {
			updatedReport.Archived = true
		} else {
			updatedReport.Archived = false
		}
		updateResult, err := reportsCollection.ReplaceOne(ctx, filter, updatedReport)
		if err != nil || updateResult.MatchedCount == 0 {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(updatedReport)
	} else {
		w.WriteHeader(403)
		return
	}
}
func auth(w http.ResponseWriter, r *http.Request) {
	var authResponse models.AuthResponse
	var user models.User
	var loggingUser models.UserCredentials

	hash,_ := bcrypt.GenerateFromPassword([]byte("cheese"), bcrypt.MinCost)
	log.Info(string(hash))
	json.NewDecoder(r.Body).Decode(&loggingUser)
	filter := bson.M{"username" : loggingUser.Username}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := usersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No such user"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loggingUser.Password))
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Password is incorrect"))
		return
	}
	tokenString := utils.CreateToken(user,w)
	authResponse.AccessToken = tokenString
	authResponse.ID = user.ID.Hex()
	authResponse.Username = user.Username
	authResponse.Name = user.Name
	authResponse.Groups = user.Groups
	authResponse.Role = user.Role
	json.NewEncoder(w).Encode(authResponse)
}