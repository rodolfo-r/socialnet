package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/techmexdev/handlertest"
	"github.com/techmexdev/the_social_network/pkg/handler"
	"github.com/techmexdev/the_social_network/pkg/storage/mock"
)

func TestSettingsInvalidCreds(t *testing.T) {
	t.Parallel()
	r := handler.New(mock.New(), handler.Options{})
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
			handlertest.Test(t, tcs[i], r)
		})
	}
}

func TestSettingsValidCreds(t *testing.T) {
	r := handler.New(mock.New(), handler.Options{})

	token, err := getAuthToken(r)
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
			handlertest.Test(t, tcs[i], r)
		})
	}
}
