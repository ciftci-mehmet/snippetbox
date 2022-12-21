package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {

	app := newTestApplication(t)

	// create new test server
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	// check value of response status code and body
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
