package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/techmexdev/socialnet"
)

// Feed sends the 'socialnet_token' as an auth
// token to the api server, and responds with an html template.
// responds with a template with the response data.
func Feed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/feed/index.html")
	if err != nil {
		serverError(w, err)
		return
	}

	apiReq, err := http.NewRequest("GET", "http://localhost:3001/feed", nil)
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

	var feed socialnet.Feed
	err = json.NewDecoder(res.Body).Decode(&feed)
	if err != nil {
		serverError(w, err)
		return
	}

	t.Execute(w, feed)
}
