package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) {
	un := mux.Vars(r)["username"]
	p, err := h.store.GetProfile(un)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	writeJSON(w, p, http.StatusOK)
}
