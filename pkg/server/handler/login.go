package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rodolfo-r/socialnet"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var usr socialnet.User

	err := json.NewDecoder(r.Body).Decode(&usr)
	defer r.Body.Close()
	if err != nil {
		http.Error(
			w,
			`Request body must be in the format: {"username": "jlennon", "password": "5tr4wb3rryfi31d5"}`,
			http.StatusBadRequest,
		)
		return
	}

	err = h.userSvc.Auth.ValidateCreds(usr.Username, usr.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	token, err := h.userSvc.Auth.CreateToken(usr.Username)
	if err != nil {
		serverError(w, err)
		return
	}

	h.r.JSON(w, http.StatusOK, map[string]string{"token": token})
}
