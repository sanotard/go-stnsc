package stnsc

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/STNS/STNS/stns"
)

const (
	successUName      = "foo"
	successUName2     = "bar"
	successUid        = 1001
	successUid2       = 1002
	successPass       = "test"
	successSha512Pass = "$6$72qH5tfJta43J1lH$o1OvvIxkDCNZtrAh3UWM9dKkGawTuBeGpLoxRuICH6B/9.Y5PA/bDtvm.fK/bB8zFNNofus6jQHXzMyiqCCqj0"
	failedUName       = "aaaaa"
	failedPass        = "aaaaa"
	failedUid         = 9999
)

var (
	successMeta = &stns.MetaData{
		ApiVersion: 2.1,
		Result:     "success",
		MinId:      successUid,
	}
	successUser = &stns.User{
		Password:  successSha512Pass,
		GroupId:   successUid,
		Directory: "/home/" + successUName,
		Shell:     "/bin/bash",
		Gecos:     "description",
		Keys:      []string{"key"},
		LinkUsers: []string{"linkuer"},
	}
	successUser2 = &stns.User{
		Password:  successSha512Pass,
		GroupId:   successUid2,
		Directory: "/home/" + successUName2,
		Shell:     "/bin/bash",
		Gecos:     "description",
		Keys:      []string{"key"},
		LinkUsers: []string{"linkuer"},
	}

	userGetItems          stns.Attributes
	successUserGetResp    *stns.ResponseFormat
	successUserGetHandler func(w http.ResponseWriter, r *http.Request)

	userListItems          stns.Attributes
	successUserListResp    *stns.ResponseFormat
	successUserListHandler func(w http.ResponseWriter, r *http.Request)
)

func init() {
	// handler for list
	userListItems = make(stns.Attributes)
	userListItems[successUName] = &stns.Attribute{
		Id:   successUid,
		User: successUser,
	}
	userListItems[successUName2] = &stns.Attribute{
		Id:   successUid2,
		User: successUser2,
	}

	successUserListResp = &stns.ResponseFormat{
		MetaData: successMeta,
		Items:    userListItems,
	}

	successUserListHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, marshal(successUserListResp))
	}

	// handler for get
	userGetItems = make(stns.Attributes)
	userGetItems[successUName] = &stns.Attribute{
		Id:   successUid,
		User: successUser,
	}
	successUserGetResp = &stns.ResponseFormat{
		MetaData: successMeta,
		Items:    userGetItems,
	}

	successUserGetHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, marshal(successUserGetResp))
	}
}

func TestUserList(t *testing.T) {
	Setup(nil)
	defer Teardown()

	// List
	mux.HandleFunc("/user/list/", successUserListHandler)

	// test response data
	var err error
	var want stns.Attributes
	var result stns.Attributes

	// success
	want = userListItems
	result, err = client.User.List()
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}
}

func TestUserGet(t *testing.T) {
	Setup(nil)
	defer Teardown()

	// Get
	mux.HandleFunc("/user/name/"+successUName, successUserGetHandler)
	// GetById
	mux.HandleFunc(fmt.Sprintf("/user/id/%v", successUid), successUserGetHandler)

	// test response data
	var err error
	var want *stns.Attribute
	var result *stns.Attribute

	// success
	want = userGetItems[successUName]
	result, err = client.User.Get(successUName)
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// failed
	want = nil
	result, err = client.User.Get(failedUName)
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// success
	want = userGetItems[successUName]
	result, err = client.User.GetById(successUid)
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// failed
	want = nil
	result, err = client.User.GetById(failedUid)
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}
}

func TestUserAuth(t *testing.T) {
	Setup(nil)
	defer Teardown()

	// Authenticate
	mux.HandleFunc("/user/name/"+successUName, successUserGetHandler)

	// test response data
	var err error
	var want *stns.Attribute
	var result *stns.Attribute

	//
	want = userListItems[successUName]
	result, err = client.User.Authenticate(successUName, successPass)
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// passowrd failed
	want = nil
	result, err = client.User.Authenticate(successUName, failedPass)
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// user not found
	want = nil
	result, err = client.User.Authenticate(failedUName, successPass)
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}
}
