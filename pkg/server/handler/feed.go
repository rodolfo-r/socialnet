package handler

import (
	"net/http"
	"strings"

	"github.com/techmexdev/socialnet"
)

func (h *handler) Feed(w http.ResponseWriter, r *http.Request) {
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	username, err := h.userSvc.Auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token. Header 'Authorization' must have value: 'Bearer <token>'", http.StatusUnauthorized)
		return
	}

	fols, err := h.userSvc.Follow.Following(username)
	if err != nil {
		serverError(w, err)
		return
	}

	var feed socialnet.Feed

	for _, f := range fols {
		usr, err := h.userSvc.Store.Read(f.Username)
		if err != nil {
			serverError(w, err)
			return
		}

		for _, p := range usr.Posts {
			feed = append(feed, socialnet.FeedItem{
				ProfileImageURL: usr.ImageURL, Post: p,
			})
		}
	}

	h.r.JSON(w, http.StatusOK, feed)
}
