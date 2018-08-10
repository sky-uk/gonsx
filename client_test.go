package gonsx

import (
	"encoding/base64"
	"fmt"
	"github.com/tadaweb/gonsx/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var user = "nsxUser"
var password = "nsxPass"
var ignoreSSL = true
var debug = false
var nsxClient *NSXClient

var server *httptest.Server

const (
	unauthorizedStatusCode = http.StatusForbidden
	unauthorizedResponse   = `<html><head><title>Pivotal tc Runtime 3.1.2.RELEASE/7.0.64.B.RELEASE - Error report</title><style><!--H1 {font-family:Tahoma,Arial,sans-serif;color:white;background-color:#525D76;font-size:22px;} H2 {font-family:Tahoma,Arial,sans-serif;color:white;background-color:#525D76;font-size:16px;} H3 {font-family:Tahoma,Arial,sans-serif;color:white;background-color:#525D76;font-size:14px;} BODY {font-family:Tahoma,Arial,sans-serif;color:black;background-color:white;} B {font-family:Tahoma,Arial,sans-serif;color:white;background-color:#525D76;} P {font-family:Tahoma,Arial,sans-serif;background:white;color:black;font-size:12px;}A {color : black;}A.name {color : black;}HR {color : #525D76;}--></style> </head><body><h1>HTTP Status 403 - VC user does not have any role on NSX Manager.</h1><HR size="1" noshade="noshade"><p><b>type</b> Status report</p><p><b>message</b> <u>VC user does not have any role on NSX Manager.</u></p><p><b>description</b> <u>Access to the specified resource has been forbidden.</u></p><HR size="1" noshade="noshade"><h3>Pivotal tc Runtime 3.1.2.RELEASE/7.0.64.B.RELEASE</h3></body></html>`
)

func hasHeader(req *http.Request, name string, value string) bool {
	return req.Header.Get(name) == value
}

func setup(statusCode int, responseBody string) {
	basicAuthHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	server = httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !hasHeader(r, "Authorization", basicAuthHeaderValue) {
				w.WriteHeader(unauthorizedStatusCode)
				fmt.Fprint(w, unauthorizedResponse)
				return
			}
			w.WriteHeader(statusCode)
			fmt.Fprintln(w, responseBody)
		}))
	nsxClient = NewNSXClient(server.URL, user, password, ignoreSSL, debug)
}

func TestHappyCase(t *testing.T) {
	setup(200, "pong")
	nsxClient = NewNSXClient(server.URL, user, password, ignoreSSL, debug)
	apiRequest := api.NewBaseAPI(http.MethodGet, "/", nil, nil)

	err := nsxClient.Do(apiRequest)

	assert.Nil(t, err)
}

// TODO: add TestFailWhenNotValidSSLCerts(t *testing.T)

func TestBasicAuthFailure(t *testing.T) {
	setup(0, "")
	nsxClient = NewNSXClient(server.URL, "invalidUser", "invalidPass", ignoreSSL, debug)

	apiRequest := api.NewBaseAPI(http.MethodGet, "/", nil, nil)
	nsxClient.Do(apiRequest)

	assert.Equal(t, 403, apiRequest.StatusCode())
	assert.Equal(t, unauthorizedResponse, string(apiRequest.RawResponse()))

}

func TestIsXML(t *testing.T) {
	assert.True(t, isXML("application/xml"))
	assert.True(t, isXML("text/xml"))
	assert.True(t, isXML("text/xml; charset=utf-8"))

	assert.False(t, isXML("application/html"))
	assert.False(t, isXML("text/html"))
	assert.False(t, isXML("text/html; charset=utf-8"))
}

// TODO: add invalid xml (parsing failure)
// TODO: add timeout
// TODO: add refused?
// TODO: add error response (5xx)
