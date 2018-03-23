package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	b, err := json.Marshal(data)
	if err != nil {
		serverError(w, err)
		return
	}

	w.WriteHeader(status)
	w.Write(b)
}

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	log.Println(err)
}

func createToken(username string) (string, error) {
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    address,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(signature))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
