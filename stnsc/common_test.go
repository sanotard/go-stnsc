package stnsc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

var (
	server *httptest.Server
	mux    *http.ServeMux
	client *Client
)

func Setup(httpClient *http.Client) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	var err error
	client, err = NewClient(server.URL, httpClient)
	if err != nil {
		panic(err)
	}

}

func Teardown() {
	server.Close()
}

func marshal(v interface{}) string {
	out, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(out)
}
