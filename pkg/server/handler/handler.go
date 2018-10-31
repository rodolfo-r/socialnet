package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfo-r/socialnet"
	"github.com/unrolled/render"
)

// Options is configuration for the router
type Options struct {
	Log       bool
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
		{method: "POST", path: "/signup", handler: h.SignUp},
		{method: "POST", path: "/login", handler: h.Login},
		{method: "GET", path: "/settings", handler: h.Settings},
		{method: "POST", path: "/submit-post", handler: h.SubmitPost},
		{method: "GET", path: "/user/{username}", handler: h.Profile},
		{method: "POST", path: "/profile-picture", handler: h.ProfilePicture},
		{method: "GET", path: "/users", handler: h.Users},
		{method: "POST", path: "/follow", handler: h.Follow},
		{method: "POST", path: "/unfollow", handler: h.Unfollow},
		{method: "GET", path: "/feed", handler: h.Feed},
		{method: "POST", path: "/like", handler: h.Like},
		{method: "POST", path: "/comment", handler: h.Comment},
	}

	for _, r := range rr {
		server.router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	server.router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))

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
