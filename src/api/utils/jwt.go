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
	expirationTime:= time.Now().Add(30*time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role" : user.Role,
		"name" : user.Name,
		"groups" : user.Groups,
		"iat" : expirationTime,
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.WithFields(log.Fields{
			"function" : "token.SignedString",
			"error"	:	err,
		},
		).Warn("Failed to create signed token!")
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return tokenString
}