package handler

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/techmexdev/the_social_network/pkg/storage"
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
	store storage.Storage
}

var signature string
var address string

// New creates a router with all handlers
func New(store storage.Storage, options Options) *mux.Router {
	router := mux.NewRouter()
	h := handler{store}
	routes := []route{
		{method: "POST", path: "/signup", handler: h.SignUp},
		{method: "POST", path: "/login", handler: h.Login},
		{method: "GET", path: "/settings", handler: AuthMiddleware(h.Settings)},
		{method: "POST", path: "/submit-post", handler: AuthMiddleware(h.SubmitPost)},
		{method: "GET", path: "/{username}", handler: h.Profile},
	}

	for _, r := range routes {
		router.Methods(r.method).Path(r.path).HandlerFunc(r.handler)
	}

	signature = options.Signature
	address = options.Address
	return router
}

// JWT config

// Claims describe jwt format
type Claims struct {
	Username string `json:"usn"`
	jwt.StandardClaims
}

type ctxKey string

const ctxUsnKey ctxKey = "username"
