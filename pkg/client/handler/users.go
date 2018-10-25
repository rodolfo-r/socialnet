package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodolfo-r/socialnet"
)

// Users responds with a template containing all users
// in the application.
func (h *handler) Users(w http.ResponseWriter, r *http.Request) {
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

	var uis []socialnet.UserItem
	err = json.NewDecoder(res.Body).Decode(&uis)
	if err != nil {
		serverError(w, err)
		return
	}

	err = h.template.ExecuteTemplate(w, "users.html", uis)
	if err != nil {
		serverError(w, fmt.Errorf("failed to execute template users.html: %s", err))
		return
	}
}
