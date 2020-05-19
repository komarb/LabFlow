package utils

import (
	"LabFlow/models"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func CreateToken(user models.User,w http.ResponseWriter) string {
	hmacSecret := []byte("test")
	iatTime := time.Now().Unix()
	expirationTime:= time.Now().Add(30*time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role" : user.Role,
		"name" : user.Name,
		"groups" : user.Groups,
		"iat" : iatTime,
		"exp" : expirationTime,
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "token.SignedString",
			"error"	:	err,
		},
		).Warn("Failed to create signed token!")
	}

	return tokenString
}