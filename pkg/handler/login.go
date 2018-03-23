package handler

import (
	"encoding/json"
	"net/http"

	"github.com/techmexdev/the_social_network/pkg/model"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var usr model.User

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, `Request body must be in the format {"username": "jlennon", "password": "5tr4wb3rryfi31d5"`, http.StatusBadRequest)
		return
	}

	err = h.store.ValidateUserCreds(usr.Username, usr.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	token, err := createToken(usr.Username)
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, map[string]string{"token": token}, 200)
}
