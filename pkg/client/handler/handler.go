package handler

import (
	"net/http"

	"html/template"

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

type handler struct {
	client   *http.Client
	template *template.Template
}

// New creates a server that responds to http requests.
func New() *Server {
	server := &Server{router: mux.NewRouter()}

	t := template.Must(template.ParseFiles(genTemplatePaths()...))

	h := handler{template: t, client: http.DefaultClient}
	rr := []route{
		{method: "GET", path: "/", handler: h.Home},
		{method: "GET", path: "/settings", handler: h.Settings},
		{method: "GET", path: "/user/{username}", handler: h.Profile},
		{method: "GET", path: "/users", handler: h.Users},
		{method: "GET", path: "/feed", handler: h.Feed},
	}

	server.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	for _, r := range rr {
		server.router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	return server
}

// ServeHTTP responds to http requests by delegating to the server.router.
func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func genTemplatePaths() []string {
	var sharedPaths, notSharedPaths []string

	sharedFilenames := []string{"header", "head", "footer"}
	for _, fn := range sharedFilenames {
		sharedPaths = append(sharedPaths, genTemplatePath(fn, true))
	}

	notSharedFilenames := []string{"index", "profile", "feed", "settings", "users"}
	for _, fn := range notSharedFilenames {
		notSharedPaths = append(notSharedPaths, genTemplatePath(fn, false))
	}

	return append(sharedPaths, notSharedPaths...)
}

func genTemplatePath(filename string, shared bool) string {
	if shared {
		return "./static/_shared/" + filename + ".html"
	}
	return "./static/" + filename + "/" + filename + ".html"
}
