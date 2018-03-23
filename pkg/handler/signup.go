package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/techmexdev/the_social_network/pkg/model"
)

// Signup creates user in storage, and responds with auth token
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var usr model.User

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w,
			`Request Body must be in the format: {"username": "jl", "firstName": "John", "lastName": "lennon", "email" "strawberry@fields.com", "password": "berrystraw123"}`,
			http.StatusBadRequest)
		return
	}

	if err := usr.Validate(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = h.store.GetUser(usr)
	if err == nil {
		log.Print(err)
		http.Error(w, "User with same username/email already exists", http.StatusBadRequest)
		return
	}

	newUsr, err := h.store.CreateUser(usr)
	if err != nil {
		serverError(w, err)
		return
	}

	token, err := createToken(newUsr.Username)
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, map[string]string{"token": token}, 201)
}
