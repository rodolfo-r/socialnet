package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/techmexdev/socialnet"
)

// Feed sends the 'socialnet_token' as an auth
// token to the api server, and responds with an html template.
// responds with a template with the response data.
func (h *handler) Feed(w http.ResponseWriter, r *http.Request) {
	apiReq, err := http.NewRequest("GET", "http://localhost:3001/feed", nil)
	if err != nil {
		serverError(w, fmt.Errorf("failed creating GET localhost:3001/feed request: %s", err))
		return
	}

	tokenCookie, err := r.Cookie("socialnet_token")
	if err != nil {
		serverError(w, fmt.Errorf("failed reading social_net cookie from request: %s", err))
		return
	}

	apiReq.Header.Set("Authorization", "Bearer "+tokenCookie.Value)

	res, err := h.client.Do(apiReq)
	if err != nil {
		serverError(w, fmt.Errorf("failed to decode api server res body: %s", err))
		return
	}

	var feed []socialnet.FeedItem
	err = json.NewDecoder(res.Body).Decode(&feed)
	if err != nil {
		serverError(w, fmt.Errorf("failed decoding feed from server api response: %s", err))
		return
	}

	err = h.template.ExecuteTemplate(w, "feed.html", feed)
	if err != nil {
		serverError(w, fmt.Errorf("failed to execute template feed.html: %s", err))
		return
	}

}
