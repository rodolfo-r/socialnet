package handler

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techmexdev/socialnet"
)

// Profile sends the 'socialnet_token' as an auth
// token to the api server, and responds with an html template.
// responds with a template with the response data.
func Profile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/profile/index.html")
	if err != nil {
		serverError(w, err)
		return
	}

	username := mux.Vars(r)["username"]
	apiReq, err := http.NewRequest("GET", "http://localhost:3001/user/"+username, nil)
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

	if res.StatusCode != 200 {
		var b []byte
		_, err := res.Body.Read(b)
		if err != nil {
			serverError(w, err)
			return
		}
		serverError(w, errors.New(string(b)))
	}
	var prof socialnet.Profile
	err = json.NewDecoder(res.Body).Decode(&prof)
	if err != nil {
		serverError(w, err)
		return
	}

	t.Execute(w, prof)
}
