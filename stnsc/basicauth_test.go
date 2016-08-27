package stnsc

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/STNS/STNS/stns"
)

var (
	basicUsername    string
	basicPassowrd    string
	basicAuthHandler func(w http.ResponseWriter, r *http.Request)
)

func checkAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return username == basicUsername && password == basicPassowrd
}

func init() {
	basicUsername = "username"
	basicPassowrd = "password"
	basicAuthHandler = func(w http.ResponseWriter, r *http.Request) {
		if !checkAuth(r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, marshal(successUserListResp))
	}
}

func TestBasicAuth(t *testing.T) {
	// failed
	Setup(nil)
	mux.HandleFunc("/user/list", basicAuthHandler)

	// test response data
	var err error
	var want stns.Attributes
	var result stns.Attributes

	want = nil
	result, err = client.User.List()
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
		t.Errorf("\ngot %v\nwant %v\n", result, nil)
	}
	Teardown()

	// success
	tp := &BasicAuthTransport{
		Username: basicUsername,
		Password: basicPassowrd,
	}
	Setup(tp.Client())
	mux.HandleFunc("/user/list", basicAuthHandler)

	want = userListItems
	result, err = client.User.List()
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

}
