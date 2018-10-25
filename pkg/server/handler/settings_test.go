package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/rodolfo-r/handlertest"
	"github.com/rodolfo-r/socialnet"
	"github.com/rodolfo-r/socialnet/pkg/server/auth"
	"github.com/rodolfo-r/socialnet/pkg/server/handler"
	"github.com/rodolfo-r/socialnet/pkg/server/storage/memo"
)

func TestSettingsInvalidCreds(t *testing.T) {
	t.Parallel()

	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	tcs := []handlertest.TestCase{
		{
			Name:       "No credentials",
			StatusCode: http.StatusUnauthorized,
			Request:    httptest.NewRequest("GET", "/settings", nil),
		},
		{
			Name:       "Bad Credentials",
			StatusCode: http.StatusUnauthorized,
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/settings"},
				Header: map[string][]string{"Authorization": {"Bearer " +
					"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o",
				}},
				Body: ioutil.NopCloser(strings.NewReader("")),
			},
		},
	}
	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}

func TestSettingsValidCreds(t *testing.T) {
	usrStore := memo.NewUserStorage()
	usrSvc := socialnet.UserService{Store: usrStore, Auth: *auth.New(usrStore, "localhost", "test jwt string")}
	postSvc := socialnet.PostService{Store: memo.NewPostStorage()}

	router := handler.New(usrSvc, postSvc, handler.Options{
		Signature: "jwt test signature",
	})

	token, err := getAuthToken(router)
	if err != nil {
		t.Fatal(err)
	}

	tcs := []handlertest.TestCase{
		{
			Name:       "Valid Credentials",
			StatusCode: http.StatusOK,
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "/settings"},
				Header: map[string][]string{"Authorization": {"Bearer " + token}},
				Body:   ioutil.NopCloser(strings.NewReader("")),
			},
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), `"email"`), "res body does not contain email")
			},
		},
	}

	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}
