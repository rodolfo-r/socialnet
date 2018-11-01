package handler

import (
	"html/template"
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
