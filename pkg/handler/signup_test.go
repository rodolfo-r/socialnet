package handler_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/techmexdev/handlertest"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/auth"
	"github.com/techmexdev/socialnet/pkg/storage/memo"
	"github.com/techmexdev/socialnet/pkg/handler"
)

func TestSignUp(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore)}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tcs := []handlertest.TestCase{
		{
			Name:       "New User",
			StatusCode: 201,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/signup"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{"username": "jlennon", "firstName": "John", "lastName": "Lennon", "email": "jlennon@beatles.com", "password": "str4wb3rryfi31d5"}`,
				)),
			},
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), `"token"`), fmt.Sprintf("expected token field in response body"))
			},
		},
		{
			Name:       "Username/email already taken",
			StatusCode: 400,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/signup"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{"username": "jlennon", "firstName": "John", "lastName": "Lennon", "email": "jlennon@beatles.com", "password": "str4wb3rryfi31d5"}`,
				)),
			},
		},
		{
			Name:       "Malformed json",
			StatusCode: 400,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/signup"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{username: "jlennon", firstName: "John", lastName: "Lennon", email: "jlennon@beatles.com", password: "str4wb3rryfi31d5"}`,
				)),
			},
		},
		{
			Name:       "Missing fields",
			StatusCode: 400,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/signup"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body: ioutil.NopCloser(strings.NewReader(
					`{username: "jlennon", email: "jlennon@beatles.com", password: "str4wb3rryfi31d5"}`,
				)),
			},
		},
	}

	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}
