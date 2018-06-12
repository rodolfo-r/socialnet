package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type follow struct {
	Username string `json:"username"`
}

func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	var fol follow
	err = json.NewDecoder(r.Body).Decode(&fol)
	if err != nil {
		serverError(w, err)
		return
	}

	err = h.userSvc.Follow.Follow(username, fol.Username)
	if err != nil {
		serverError(w, err)
		return
	}

	h.r.JSON(w, http.StatusOK, http.StatusText(http.StatusOK))
}
