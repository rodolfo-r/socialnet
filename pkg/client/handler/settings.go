package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rodolfo-r/socialnet"
)

// Settings sends the 'socialnet_token' as an auth
// token to the api server, and responds with an html template.
func (h *handler) Settings(w http.ResponseWriter, r *http.Request) {
	apiReq, err := http.NewRequest("GET", "http://localhost:3001/settings", nil)
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
	defer res.Body.Close()

	var set socialnet.Settings
	err = json.NewDecoder(res.Body).Decode(&set)
	if err != nil {
		serverError(w, err)
		return
	}

	err = h.template.ExecuteTemplate(w, "settings.html", set)
	if err != nil {
		serverError(w, err)
		return
	}
}
