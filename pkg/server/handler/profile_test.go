package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/techmexdev/handlertest"
	"github.com/techmexdev/socialnet"
	"github.com/techmexdev/socialnet/pkg/server/auth"
	"github.com/techmexdev/socialnet/pkg/server/handler"
	"github.com/techmexdev/socialnet/pkg/server/storage/memo"
)

func TestProfile(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tcs := []handlertest.TestCase{
		{
			Name:       "User doesn't exist",
			StatusCode: http.StatusNotFound,
			Request:    httptest.NewRequest("GET", "/jlennon", nil),
		},
		{
			Name: "Signup new user",
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
			Name:       "User exists",
			Request:    httptest.NewRequest("GET", "/user/jlennon", nil),
			StatusCode: http.StatusOK,
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), "firstName"), "should contain firstName property")
			},
		},
	}
	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}
