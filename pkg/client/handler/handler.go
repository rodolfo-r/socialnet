package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server responds to http requests.
type Server struct {
	router *mux.Router
}

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

// New creates a server that responds to http requests.
func New() *Server {
	server := &Server{router: mux.NewRouter()}

	rr := []route{
		{method: "GET", path: "/", handler: Home},
		{method: "GET", path: "/settings", handler: Settings},
		{method: "GET", path: "/user/{username}", handler: Profile},
		{method: "GET", path: "/users", handler: Users},
	}

	server.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	for _, r := range rr {
		server.router.HandleFunc(r.path, r.handler).Methods(r.method)
	}
	server.router.Use(LogMiddleware)

	return server
}

// ServeHTTP responds to http requests by delegating to the server.router.
func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

// Home responds with the root index.html.
func Home(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static/index/")).ServeHTTP(w, r)
}
