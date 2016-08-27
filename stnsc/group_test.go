package stnsc

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/STNS/STNS/stns"
)

const (
	successGName  = "foo"
	successGName2 = "bar"
	successGid    = 1002
	successGid2   = 1002
	failedGName   = "aaaaa"
	failedGid     = 9999
)

var (
	successGroup = &stns.Group{
		Users:      []string{"user1", "user2"},
		LinkGroups: []string{"linkuer1", "linkuser2"},
	}
	successGroup2 = &stns.Group{
		Users:      []string{"user11", "user22"},
		LinkGroups: []string{"linkuer11", "linkuser22"},
	}

	groupGetItems          stns.Attributes
	successGroupGetResp    *stns.ResponseFormat
	successGroupGetHandler func(w http.ResponseWriter, r *http.Request)

	groupListItems          stns.Attributes
	successGroupListResp    *stns.ResponseFormat
	successGroupListHandler func(w http.ResponseWriter, r *http.Request)
)

func init() {
	// List
	groupListItems = make(stns.Attributes)
	groupListItems[successGName] = &stns.Attribute{
		Id:    successGid,
		Group: successGroup,
	}
	groupListItems[successGName2] = &stns.Attribute{
		Id:    successGid2,
		Group: successGroup2,
	}

	successGroupListResp = &stns.ResponseFormat{
		MetaData: successMeta,
		Items:    groupListItems,
	}

	successGroupListHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, marshal(successGroupListResp))
	}

	// Get
	groupGetItems = make(stns.Attributes)
	groupGetItems[successGName] = &stns.Attribute{
		Id:    successGid,
		Group: successGroup,
	}
	successGroupGetResp = &stns.ResponseFormat{
		MetaData: successMeta,
		Items:    groupGetItems,
	}

	successGroupGetHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, marshal(successGroupGetResp))
	}
}

func TestGroupList(t *testing.T) {
	Setup(nil)
	defer Teardown()

	// List
	mux.HandleFunc("/group/list/", successGroupListHandler)

	// test response data
	var err error
	var want stns.Attributes
	var result stns.Attributes

	// success
	want = groupListItems
	result, err = client.Group.List()
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}
}

func TestGroupGet(t *testing.T) {
	Setup(nil)
	defer Teardown()

	// Get
	mux.HandleFunc("/group/name/"+successGName, successGroupGetHandler)
	// GetById
	mux.HandleFunc(fmt.Sprintf("/group/id/%v", successGid), successGroupGetHandler)

	// test response data
	var err error
	var want *stns.Attribute
	var result *stns.Attribute

	// success
	want = groupGetItems[successGName]
	result, err = client.Group.Get(successGName)
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// failed
	want = nil
	result, err = client.Group.Get(failedGName)
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// success
	want = groupGetItems[successGName]
	result, err = client.Group.GetById(successGid)
	if err != nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}

	// failed
	want = nil
	result, err = client.Group.GetById(failedGid)
	if err == nil {
		t.Errorf("\ngot %v\nwant %v\n", err, nil)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("\ngot %#v\n\nwant %#v\n", result, want)
	}
}
