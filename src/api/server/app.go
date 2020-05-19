package server

import (
	"LabFlow/config"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	DB *mongo.Client
}

var subjectsCollection *mongo.Collection
var usersCollection *mongo.Collection
var tasksCollection *mongo.Collection
var reportsCollection *mongo.Collection
var groupsCollection *mongo.Collection
var cfg *config.Config

func (a *App) Init(config *config.Config) {
	cfg = config
	fmt.Println(cfg.DB.DBPort)
	DBUri := "mongodb://" + cfg.DB.Host + ":" + cfg.DB.DBPort
	log.WithField("dburi", DBUri).Info("Current database URI: ")
	client, err := mongo.NewClient(options.Client().ApplyURI(DBUri))
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.NewClient",
			"error"	:	err,
			"db_uri":	DBUri,
		},
		).Fatal("Failed to create new MongoDB client")
	}

	// Create db connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Connect",
			"error"	:	err},
		).Fatal("Failed to connect to MongoDB")
	}

	// Check the connection
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "mongo.Ping",
			"error"	:	err},
		).Fatal("Failed to ping MongoDB")
	}
	log.Info("Connected to MongoDB!")
	log.WithFields(log.Fields{
		"db_name" : cfg.DB.DBName,
		"collection_name" : cfg.DB.SubjectsCollectionName,
	}).Info("Database information: ")
	log.WithField("testMode", cfg.App.TestMode).Info("Let's check if test mode is on...")

	subjectsCollection = client.Database(cfg.DB.DBName).Collection(cfg.DB.SubjectsCollectionName)
	usersCollection = client.Database(cfg.DB.DBName).Collection(cfg.DB.UsersCollectionName)
	tasksCollection = client.Database(cfg.DB.DBName).Collection(cfg.DB.TasksCollectionName)
	reportsCollection = client.Database(cfg.DB.DBName).Collection(cfg.DB.ReportsCollectionName)
	groupsCollection = client.Database(cfg.DB.DBName).Collection(cfg.DB.GroupsCollectionName)

	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {

	public := a.Router.PathPrefix("/auth").Subrouter()
	public.Use(loggingMiddleware)
	public.HandleFunc("", auth).Methods("POST")

	private := a.Router.PathPrefix("/api").Subrouter()
	private.Use(testAuthMiddleware)
	private.HandleFunc("/subjects", getSubjects).Methods("GET")
	private.HandleFunc("/subjects/{id}", getSubject).Methods("GET")
	private.HandleFunc("/subjects/{subjectID}", createTask).Methods("POST")
	private.HandleFunc("/subjects/{subjectID}/tasks", getSubjectTasks).Methods("GET")
	private.HandleFunc("/subjects/{subjectID}/groups", getSubjectGroups).Methods("GET")
	private.HandleFunc("/subjects/{subjectID}/tasks/{taskID}/reports", getTaskReports).Methods("GET")
	private.HandleFunc("/subjects/{subjectID}/tasks/{taskID}", createReport).Methods("POST")
	private.HandleFunc("/subjects/{subjectID}/tasks/{taskID}", createReport).Methods("OPTIONS")
	private.HandleFunc("/subjects/{subjectID}/reports/{studentID}", getStudentReports).Methods("GET")
	private.HandleFunc("/reports/{reportID}", updateReport).Methods("PUT")

	//a.Router.Use(testAuthMiddleware)


	/*a.Router.HandleFunc("/api/reports", getAllReportsSorted).Methods("GET").Queries("sorted_by","{var}")
	a.Router.HandleFunc("/api/employee/{employee}", getEmployeeReports).Methods("GET").Queries("dateBegin","{dateBegin}", "dateEnd", "{dateEnd}")
	a.Router.HandleFunc("/api/employee/{employee}", getEmployeeReports).Methods("GET")
	a.Router.HandleFunc("/api/reports", getAllReports).Methods("GET")
	a.Router.HandleFunc("/api/reports/archived", getArchivedReports).Methods("GET")
	a.Router.HandleFunc("/api/reports/{id}", getReport).Methods("GET")
	a.Router.HandleFunc("/api/reports", createReport).Methods("POST")
	a.Router.HandleFunc("/api/reports/{id}", updateReport).Methods("PUT")
	a.Router.HandleFunc("/api/reports/{id}", deleteReport).Methods("DELETE")*/



}

func (a *App) Run(addr string) {
	log.WithField("port", cfg.App.AppPort).Info("Starting server on port:")
	log.Info("\n\nNow handling routes!")

	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "http.ListenAndServe",
			"error"	:	err},
		).Fatal("Failed to run a server!")
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}