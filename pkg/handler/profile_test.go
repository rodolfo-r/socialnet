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

func TestProfile(t *testing.T) {
	t.Parallel()
	r := handler.New(mock.New(), handler.Options{})
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
			Request:    httptest.NewRequest("GET", "/jlennon", nil),
			StatusCode: http.StatusOK,
			BodyAssert: func(b []byte) error {
				return handlertest.Assert(strings.Contains(string(b), "firstName"), "should contain firstName property")
			},
		},
	}
	for i, _ := range tcs {
		t.Run(tcs[i].Name, func(t *testing.T) {
			handlertest.Test(t, tcs[i], r)
		})
	}
}
