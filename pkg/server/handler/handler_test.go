package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
)

func getAuthToken(r http.Handler) (string, error) {
	reqbody := `{"username": "jlennon", "firstName": "John", "lastName": "Lennon", "email": "jlennon@beatles.com", "password": "str4wb3rryfi31d5"}`
	req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqbody))
	if err != nil {
		return "", err
	}

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	res := rec.Result()

	token := struct {
		Value string `json:"token"`
	}{}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	return token.Value, nil
}
