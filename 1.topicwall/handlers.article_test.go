// handlers.article_tesst.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the home page returns the home page
// with the HTTP code 200(OK) for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOk := w.Code == http.StatusOK

		// Test that the pahe title is "Home Page"
		// You can carry out a lot more detailed test using libraries that can
		// parse and process HTMP pages
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Home Page</title") > 0
		return statusOk && pageOk
	})
}
