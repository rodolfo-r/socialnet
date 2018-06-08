package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/techmexdev/socialnet"
)

// Settings sends the 'socialnet_token' as an auth
// token to the api server, and responds with an html template.
func Settings(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/settings/index.html")
	if err != nil {
		serverError(w, err)
		return
	}

	apiReq, err := http.NewRequest("GET", "http://localhost:3001/api/settings", nil)
	if err != nil {
		serverError(w, err)
		return
	}

	tokenCookie, err := r.Cookie("socialnet_token")
	if err != nil {
		serverError(w, err)
		return
	}

	apiReq.Header.Set("Authorization", "Bearer "+tokenCookie.Value)

	client := &http.Client{}
	res, err := client.Do(apiReq)
	if err != nil {
		serverError(w, err)
		return
	}

	var set socialnet.Settings
	err = json.NewDecoder(res.Body).Decode(&set)
	if err != nil {
		serverError(w, err)
		return
	}

	t.Execute(w, set)
}
