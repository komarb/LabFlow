package server

import (
	"LabFlow/logging"
	"LabFlow/models"
	"github.com/auth0-community/go-auth0"
	log "github.com/sirupsen/logrus"
	"net/http"
)
var validator *auth0.JWTValidator
var Claims	models.Claims

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sw := logging.NewStatusWriter(w)
		sw.Header().Set("Content-Type", "application/json")
		sw.Header().Set("Access-Control-Allow-Origin", "*")
		sw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		sw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			sw.WriteHeader(200)
			return
		}
		next.ServeHTTP(sw, r)
		logging.LogHandler(sw, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "OPTIONS" {
			secret := []byte("test")
			secretProvider := auth0.NewKeyProvider(secret)
			configuration := auth0.NewConfigurationTrustProvider(secretProvider, nil, "")
			validator = auth0.NewValidator(configuration, nil)
			_, err := validator.ValidateRequest(r)

			if err != nil {
				log.WithFields(log.Fields{
					"requiredAlgorithm" : "HS256",
					"error" : err,
				}).Warning("Token is not valid!")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Token is not valid\nError: "))
				w.Write([]byte(err.Error()))

				return
			}
			getClaims(r)
		}
		sw := logging.NewStatusWriter(w)
		sw.Header().Set("Content-Type", "application/json")
		sw.Header().Set("Access-Control-Allow-Origin", "*")
		sw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		sw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			sw.WriteHeader(200)
			return
		}
		next.ServeHTTP(sw, r)
		logging.LogHandler(sw, r)
	})
}

func getClaims(r *http.Request) error {
	token, err := validator.ValidateRequest(r)
	if err != nil {
		return err
	}
	err = validator.Claims(r, token, &Claims)

	return nil
}

func isUser() bool {
	if Claims.Role == "student" {
		return true
	}
	return false
}

func isTeacher() bool {
	if Claims.Role == "teacher" {
		return true
	}
	return false
}