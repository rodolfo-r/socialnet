package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rodolfo-r/socialnet"
)

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) {
	un := mux.Vars(r)["username"]
	log.Println("un: ", un)
	usr, err := h.userSvc.Store.Read(un)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}

	log.Printf("usr = %+v\n", usr)
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
	log.Printf("followers = %+v\n", followers)

	following, err := h.userSvc.Follow.Following(usr.Username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	log.Printf("following = %+v\n", following)

	prof.Followers = followers
	prof.Following = following

	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	loggedInUsr, _ := h.userSvc.Auth.ValidateToken(token)
	log.Printf("loggedInUsr = %+v\n", loggedInUsr)

	for _, f := range followers {
		if f.Username == loggedInUsr {
			prof.IsFollower = true
		}
	}

	h.r.JSON(w, http.StatusOK, prof)
}
