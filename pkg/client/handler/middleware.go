package handler

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// LogMiddleware logs request proto, method, url, headers, and body
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		r.Body = ioutil.NopCloser(&buf)
		reqBody, _ := ioutil.ReadAll(tee)
		if strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "multipart/form-data") {
			reqBody = nil
		}
		log.Printf("%s - %s - %s - %s - %s ", r.Proto, r.Method, r.URL, r.Header, reqBody)
		next.ServeHTTP(w, r)
	})
}
