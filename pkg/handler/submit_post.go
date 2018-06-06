package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/techmexdev/socialnet"
)

func (h *handler) SubmitPost(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	var post socialnet.Post

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, `Request body must be in the form: 
		{"title": "I am the Walrus!",
		"body": "I am he as you are he as you are me, and we are all together ♫"}`, http.StatusBadRequest)
		return
	}

	post.Author = username

	createdPost, err := h.postSvc.Store.Create(post)
	if err != nil {
		serverError(w, err)
		return
	}

	h.r.JSON(w, http.StatusCreated, createdPost)
}
