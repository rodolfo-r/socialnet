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

		for i := range usr.Posts {
			ll, err := h.postSvc.Like.List(usr.Posts[i].ID)
			if err != nil {
				serverError(w, err)
				return
			}
			usr.Posts[i].Likes = ll
			fi := socialnet.FeedItem{
				ProfileImageURL: usr.ImageURL, Post: usr.Posts[i], Liked: userLikesPost(username, usr.Posts[i]),
			}
			feed = append(feed, fi)
		}
	}

	h.r.JSON(w, http.StatusOK, feed)
}

func userLikesPost(username string, post socialnet.Post) bool {
	var liked bool
	for _, l := range post.Likes {
		if l.Username == username {
			liked = true
		}
	}
	return liked
}
