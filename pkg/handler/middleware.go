package handler

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware calls next handler with username in context
// if auth token is valid or responds with Forbidden if not
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	authErrMsg := `Header should be in the format {Authorization: "Bearer <token>"}`
	return func(w http.ResponseWriter, r *http.Request) {
		authHead := r.Header.Get("Authorization")
		if !strings.Contains(authHead, "Bearer ") {
			http.Error(w, authErrMsg, http.StatusUnauthorized)
			return
		}
		tknStr := strings.Split(authHead, "Bearer ")[1]

		token, err := jwt.ParseWithClaims(tknStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(signature), nil
		})
		if err != nil {
			http.Error(w, authErrMsg, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}
		if claims.Username == "" {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)	
			return
		}
		ctx := context.WithValue(r.Context(), ctxUsnKey, claims.Username)
		next(w, r.WithContext(ctx))
	}
}

// LogMiddleware logs request proto, method, url, headers, and body
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		r.Body = ioutil.NopCloser(&buf)
		reqBody, _ := ioutil.ReadAll(tee)
		log.Printf("%s - %s - %s - %s - %s ", r.Proto, r.Method, r.URL, r.Header, reqBody)
		next.ServeHTTP(w, r)
	}
}
