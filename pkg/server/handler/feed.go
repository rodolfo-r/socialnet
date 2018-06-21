package handler

import (
	"net/http"
	"sort"
	"strings"

	"github.com/techmexdev/socialnet"
)

func (h *handler) Feed(w http.ResponseWriter, r *http.Request) {
	var f feed

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

	usr, err := h.userSvc.Store.Read(username)
	if err != nil {
		serverError(w, err)
		return
	}

	ui := socialnet.UserItem{
		Username: usr.Username, ImageURL: usr.ImageURL, FirstName: usr.FirstName, LastName: usr.LastName,
	}

	fols = append(fols, ui)
	for _, fol := range fols {
		usr, err := h.userSvc.Store.Read(fol.Username)
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

			cc, err := h.postSvc.Comment.List(usr.Posts[i].ID)
			if err != nil {
				serverError(w, err)
				return
			}
			usr.Posts[i].Comments = cc

			fi := socialnet.FeedItem{
				ProfileImageURL: usr.ImageURL, Post: usr.Posts[i], Liked: userLikesPost(username, usr.Posts[i]),
			}
			f.items = append(f.items, fi)
		}
	}

	sort.Sort(f)
	h.r.JSON(w, http.StatusOK, f.items)
}

type feed struct {
	items []socialnet.FeedItem
}

func (f feed) Len() int {
	return len(f.items)
}

func (f feed) Less(i, j int) bool {
	return f.items[i].CreatedAt.After(f.items[j].CreatedAt)
}

func (f feed) Swap(i, j int) {
	temp := f.items[i]
	f.items[i] = f.items[j]
	f.items[j] = temp
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
