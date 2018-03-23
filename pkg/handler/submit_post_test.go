package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/techmexdev/handlertest"
	"github.com/techmexdev/the_social_network/pkg/handler"
	"github.com/techmexdev/the_social_network/pkg/storage/mock"
)

func TestSubmitPostInvalidCreds(t *testing.T) {
	t.Parallel()
	r := handler.New(mock.New(), handler.Options{})
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
			handlertest.Test(t, tcs[i], r)
		})
	}
}

func TestSubmitPostValidCreds(t *testing.T) {
	t.Parallel()
	r := handler.New(mock.New(), handler.Options{})
	token, err := getAuthToken(r)
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
				Body: ioutil.NopCloser(strings.NewReader(`{"caption": "I am the Walrus!"}`)),
			},
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), `"I am the Walrus!"`), "should respond with created resource")
			},
		},
	}

	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], r)
		})
	}
}
