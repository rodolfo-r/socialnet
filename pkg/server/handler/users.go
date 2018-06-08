package handler

import (
	"net/http"

	"github.com/techmexdev/socialnet"
)

// Users responds with a list of all the users' profiles (except posts).
func (h *handler) Users(w http.ResponseWriter, r *http.Request) {
	uu, err := h.userSvc.Store.List()
	if err != nil {
		serverError(w, err)
		return
	}

	var pp []socialnet.Profile

	for _, u := range uu {
		pp = append(pp, socialnet.Profile{
			Username: u.Username, FirstName: u.FirstName, LastName: u.LastName, ImageURL: u.ImageURL,
		})
	}

	h.r.JSON(w, http.StatusOK, pp)
}
