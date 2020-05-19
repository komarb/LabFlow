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
		next.ServeHTTP(sw, r)
		logging.LogHandler(sw, r)
	})
}
/*func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: cfg.Auth.KeyURL}, nil)
		audience := cfg.Auth.Audience
		configuration := auth0.NewConfiguration(client, []string{audience}, cfg.Auth.Issuer, jose.RS256)
		validator = auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(r)
		if err != nil {
			log.WithFields(log.Fields{
				"requiredAlgorithm" : "RS256",
				"error" : err,
			}).Warning("Token is not valid!")

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token is not valid!\nError: "))
			w.Write([]byte(err.Error()))
			return
		}

		Claims.ITLab = nil
		err = getClaims(r)
		if err != nil {
			log.WithFields(log.Fields{
				"requiredClaims" : "iss, aud, sub, itlab",
				"error" : err,
			}).Warning("Invalid claims!")

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid claims!"))
			w.Write([]byte(err.Error()))
			return
		}

		if !checkScope(cfg.Auth.Scope) {
			log.WithFields(log.Fields{
				"requiredScope" : cfg.Auth.Scope,
				"error" : err,
			}).Warning("Invalid scope!")

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid scope"))
			return
		}

		if !isUser() {
			log.WithFields(log.Fields{
				"Claims.ITLab" : Claims.ITLab,
				"function" : "authMiddleware",
			}).Warning("Wrong itlab claim!")
			w.WriteHeader(403)
			w.Write([]byte("Wrong itlab claim!"))
			return
		}
		sw := logging.NewStatusWriter(w)
		next.ServeHTTP(sw, r)
		logging.LogHandler(sw, r)
	})
}*/

func testAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		sw := logging.NewStatusWriter(w)
		sw.Header().Set("Content-Type", "application/json")
		sw.Header().Set("Access-Control-Allow-Origin", "*")
		sw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		sw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
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