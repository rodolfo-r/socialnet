package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type unfollow struct {
	Username string `json:"username"`
}

func (h *handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	var unfol unfollow
	err = json.NewDecoder(r.Body).Decode(&unfol)
	if err != nil {
		serverError(w, err)
		return
	}

	err = h.userSvc.Follow.Unfollow(username, unfol.Username)
	if err != nil {
		serverError(w, err)
		return
	}

	h.r.JSON(w, http.StatusOK, http.StatusText(http.StatusOK))
}
