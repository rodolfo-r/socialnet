package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/techmexdev/socialnet"
)

// Signup creates user in storage, and responds with auth token
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var usr socialnet.User

	syntaxErr := `Request Body must be in the format: 
		{"username": "jl", "firstName": "John", "lastName": "lennon", "email" "strawberry@fields.com", "password": "berrystraw123"}`

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, syntaxErr, http.StatusBadRequest)
		return
	}

	_, err = h.userSvc.Store.Read(usr.Username)
	if err == nil {
		log.Print(err)
		http.Error(w, "User with same username already exists", http.StatusBadRequest)
		return
	}

	_, err = h.userSvc.Store.Create(usr)
	if err != nil {
		serverError(w, err)
		return
	}

	token, err := createToken(usr.Username)
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, map[string]string{"token": token}, 201)
}
