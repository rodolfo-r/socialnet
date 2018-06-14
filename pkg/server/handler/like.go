package handler

import (
	"net/http"
	"strings"

	"encoding/json"
)

type like struct {
	PostID string `json:"postID"`
	Like   bool   `json:"like"`
}

func (h *handler) Like(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	var l like
	err = json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, `Request body should be in format: {"postID":"123", "like": true}`, http.StatusUnauthorized)
		return
	}

	if l.Like {
		err = h.postSvc.Like.Create(username, l.PostID)
		if err != nil {
			serverError(w, err)
			return
		}
	} else {
		err = h.postSvc.Like.Delete(username, l.PostID)
		if err != nil {
			serverError(w, err)
			return
		}
	}

	h.r.JSON(w, http.StatusOK, http.StatusText(http.StatusOK))
}
