package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techmexdev/socialnet"
	"github.com/unrolled/render"
)

// Options is configuration for the router
type Options struct {
	Log       bool
	Address   string
	Signature string
}

// Server responds to http requests.
type Server struct {
	router *mux.Router
}

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

type handler struct {
	userSvc socialnet.UserService
	postSvc socialnet.PostService
	r       render.Render
}

// New creates a router with all handlers
func New(userSvc socialnet.UserService, postSvc socialnet.PostService, options Options) *Server {
	server := &Server{mux.NewRouter()}
	h := handler{userSvc, postSvc, *render.New()}

	rr := []route{
		{method: "POST", path: "/api/signup", handler: h.SignUp},
		{method: "POST", path: "/api/login", handler: h.Login},
		{method: "GET", path: "/api/settings", handler: h.Settings},
		{method: "POST", path: "/api/submit-post", handler: h.SubmitPost},
		{method: "GET", path: "/api/user/{username}", handler: h.Profile},
	}

	for _, r := range rr {
		server.router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	if options.Log {
		server.router.Use(LogMiddleware)
	}

	return server
}

// ServeHTTP handles responding to http requests.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	s.router.ServeHTTP(w, r)
}
