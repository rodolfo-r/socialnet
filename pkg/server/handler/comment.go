package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type comment struct {
	PostID string `json:"postID"`
	Text   string `json:"text"`
}

// Comment handles a user's request to comment on a post.
func (h *handler) Comment(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	var c comment
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, `Request body must be in the format: {"postID": "123", "text": "hello goodbye hello goodbye"}`, http.StatusBadRequest)
		return
	}

	err = h.postSvc.Comment.Create(username, c.PostID, c.Text)
	if err != nil {
		serverError(w, err)
		return
	}

	h.r.JSON(w, http.StatusCreated, http.StatusText(http.StatusCreated))
}
