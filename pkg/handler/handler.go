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
func New(userSvc socialnet.UserService, postSvc socialnet.PostService, options Options) *mux.Router {
	router := mux.NewRouter()
	h := handler{userSvc, postSvc, *render.New()}

	routes := []route{
		{method: "POST", path: "/signup", handler: h.SignUp},
		{method: "POST", path: "/login", handler: h.Login},
		{method: "GET", path: "/settings", handler: h.Settings},
		{method: "POST", path: "/submit-post", handler: h.SubmitPost},
		{method: "GET", path: "/{username}", handler: h.Profile},
	}

	for _, r := range routes {
		router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	if options.Log {
		router.Use(LogMiddleware)
	}

	return router
}
