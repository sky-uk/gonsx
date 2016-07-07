package client

import (
	"net/http"
	"net/http/httptest"
	"fmt"
)

var user = "nsxUser"
var password = "nsxPass"
var ignoreSSL = false
var debug = false
var nsxClient *NSXClient

var server *httptest.Server

func setup(statusCode int, responseBody string)  {
	// TODO: add basic auth checking?
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, responseBody)
	}))
	nsxClient = NewNSXClient(server.URL, user, password, ignoreSSL, debug)
}

// TODO: add invalid xml (parsing failure)
// TODO: add timeout
// TODO: add refused?
// TODO: add error response (5xx)
