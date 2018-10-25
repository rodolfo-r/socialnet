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

func TestSubmitPostInvalidCreds(t *testing.T) {
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
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/submit-post"},
				Header: map[string][]string{"Content-Type": {"Application/json"}},
				Body:   ioutil.NopCloser(strings.NewReader(`"caption": "Hello, world!"`)),
			},
		},
		{
			Name:       "Bad credentials",
			StatusCode: http.StatusUnauthorized,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/submit-post"},
				Header: map[string][]string{
					"Content-Type": {"Application/json"},
					"Authorization": {"Bearer " +
						"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o",
					}},
				Body: ioutil.NopCloser(strings.NewReader(`"caption": "Hello, world!"`)),
			},
		},
	}

	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}

func TestSubmitPostValidCreds(t *testing.T) {
	t.Parallel()

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
			StatusCode: http.StatusCreated,
			Request: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "/submit-post"},
				Header: map[string][]string{
					"Content-Type":  {"Application/json"},
					"Authorization": {"Bearer " + token},
				},
				Body: ioutil.NopCloser(strings.NewReader(`
					{"title": "I am the Walrus!", "body": "I am he as you are he as you are me, and we are all together ♫"}`)),
			},
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), "I am the Walrus!"), "should respond with created resource")
			},
		},
	}

	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], router)
		})
	}
}
