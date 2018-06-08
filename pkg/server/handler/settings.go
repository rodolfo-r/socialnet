package handler

import (
	"net/http"
	"strings"

	"github.com/techmexdev/socialnet"
)

func (h *handler) Settings(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	usr, err := h.userSvc.Store.Read(username)
	if err != nil {
		serverError(w, err)
		return
	}

	set := &socialnet.Settings{
		Username:  usr.Username,
		ImageURL:  usr.ImageURL,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	}
	if err != nil {
		serverError(w, err)
		return
	}

	h.r.JSON(w, http.StatusOK, set)
}
