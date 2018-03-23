package handler

import (
	"net/http"
)

func (h *handler) Settings(w http.ResponseWriter, r *http.Request) {
	usn := r.Context().Value(ctxUsnKey).(string)
	settings, err := h.store.GetUserSettings(usn)
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, settings, 200)

}
