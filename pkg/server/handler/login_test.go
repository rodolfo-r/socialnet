package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/rodolfo-r/handlertest"
	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/auth"
	"github.com/rodolfo-r/socialnet/pkg/server/handler"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tcs := []handlertest.TestCase{
		{
			Name:       "Signup user",
			StatusCode: http.StatusCreated,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/signup"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{"username": "jlennon", "firstName": "John", "lastName": "Lennon", "email": "jlennon@beatles.com", "password": "5tr4wb3rryfi31d5"}`,
				)),
			},
		},
		{
			Name:       "Login existing user",
			StatusCode: http.StatusOK,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/login"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{"username": "jlennon", "password": "5tr4wb3rryfi31d5"}`,
				)),
			},
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), `"token"`), "response body must contain an auth token")
			},
		},
		{
			Name:       "Login non-existing user",
			StatusCode: http.StatusForbidden,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/login"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{"username": "georgeh", "password": "my5w33tl1rd"}`,
				)),
			},
		},
		{
			Name:       "Bad Credentials",
			StatusCode: http.StatusForbidden,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/login"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{"username": "jlennon", "password": "password123"}`,
				)),
			},
		},
	}

	for i := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}
