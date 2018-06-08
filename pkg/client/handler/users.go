package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/techmexdev/socialnet"
)

// Users responds with a template containing all users
// in the application.
func Users(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/users/index.html")
	if err != nil {
		serverError(w, err)
		return
	}

	apiReq, err := http.NewRequest("GET", "http://localhost:3001/users", nil)
	if err != nil {
		serverError(w, err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(apiReq)
	if err != nil {
		serverError(w, err)
		return
	}
	defer res.Body.Close()

	var pp []socialnet.Profile
	err = json.NewDecoder(res.Body).Decode(&pp)
	if err != nil {
		serverError(w, err)
		return
	}

	t.Execute(w, pp)
}
