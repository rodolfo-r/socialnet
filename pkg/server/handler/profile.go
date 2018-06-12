package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/techmexdev/socialnet"
)

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) {
	un := mux.Vars(r)["username"]
	usr, err := h.userSvc.Store.Read(un)

	prof := &socialnet.Profile{
		Username:  usr.Username,
		ImageURL:  usr.ImageURL,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Posts:     usr.Posts,
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	followers, err := h.userSvc.Follow.Followers(usr.Username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	following, err := h.userSvc.Follow.Following(usr.Username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	prof.Followers = followers
	prof.Following = following

	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	loggedInUsr, _ := h.userSvc.Auth.ValidateToken(token)

	for _, f := range followers {
		if f.Username == loggedInUsr {
			prof.IsFollower = true
		}
	}

	h.r.JSON(w, http.StatusOK, prof)
}
