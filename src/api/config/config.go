package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	DB *DBConfig	`json:"DbOptions"`
	Auth *AuthConfig	`json:"AuthOptions"`
	App *AppConfig		`json:"AppOptions"`
}

type DBConfig struct {
	Host 		string		`json:"host"`
	DBPort 		string		`json:"dbPort"`
	DBName 		string		`json:"dbName"`
	SubjectsCollectionName 	string 	`json:"subjectsCollectionName"`
	UsersCollectionName		string	`json:"usersCollectionName"`
	TasksCollectionName		string	`json:"tasksCollectionName"`
	ReportsCollectionName		string	`json:"reportsCollectionName"`
	GroupsCollectionName		string	`json:"groupsCollectionName"`
}
type AuthConfig struct {
	KeyURL		string		`json:"keyUrl"`
	Audience	string		`json:"audience"`
	Issuer		string		`json:"issuer"`
	Scope		string		`json:"scope"`
}
type AppConfig struct {
	AppPort		string	`json:"appPort"`
	TestMode	bool	`json:"testMode"`
}

func GetConfig() *Config {
	var config Config
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "GetConfig.ReadFile",
			"error"	:	err,
		},
		).Fatal("Can't read config.json file, shutting down...")
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "GetConfig.Unmarshal",
			"error"	:	err,
		},
		).Fatal("Can't correctly parse json from config.json, shutting down...")
	}

	data, err = ioutil.ReadFile("auth_config.json")
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "GetConfig.ReadFile",
			"error"	:	err,
		},
		).Fatal("Can't read auth_config.json file, shutting down...")
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "GetConfig.Unmarshal",
			"error"	:	err,
		},
		).Fatal("Can't correctly parse json from auth_config.json, shutting down...")
	}
	return &config
}
