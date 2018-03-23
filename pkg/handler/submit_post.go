package handler

import (
	"encoding/json"
	"net/http"

	"github.com/techmexdev/the_social_network/pkg/model"
)

func (h *handler) SubmitPost(w http.ResponseWriter, r *http.Request) {
	un := r.Context().Value(ctxUsnKey).(string)
	cpt := struct {
		Value string `json:"caption"`
	}{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&cpt)
	if err != nil {
		http.Error(w, `Request body must be in the form: {"caption": "I am the Walrus!"`, http.StatusBadRequest)
		return
	}

	post := model.Post{
		Username: un, Caption: cpt.Value,
	}
	createdPost, err := h.store.CreatePost(post)
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, createdPost, http.StatusCreated)
}
