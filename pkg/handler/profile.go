package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techmexdev/socialnet"
)

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) {
	un := mux.Vars(r)["username"]
	usr, err := h.userSvc.Store.Read(un)

	prof := &socialnet.Profile{
		Username: usr.Username, FirstName: usr.FirstName, LastName: usr.LastName, Posts: usr.Posts,
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	h.r.JSON(w, http.StatusOK, prof)
}
