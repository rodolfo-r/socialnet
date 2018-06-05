package handler

import (
	"net/http"

	"github.com/techmexdev/socialnet"
)

func (h *handler) Settings(w http.ResponseWriter, r *http.Request) {
	un := r.Context().Value(ctxUsnKey).(string)
	usr, err := h.userSvc.Store.Read(un)

	set := &socialnet.Settings{
		Username: usr.Username, FirstName: usr.FirstName, LastName: usr.LastName, Email: usr.Email,
	}
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, set, 200)
}
