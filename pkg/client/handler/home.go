package handler

import "net/http"

// Home responds with the root index.html.
func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static/index/")).ServeHTTP(w, r)
}
