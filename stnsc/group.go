package stnsc

import (
	"errors"
	"fmt"

	"github.com/STNS/STNS/stns"
)

var (
	// Group not found
	ErrGroupNotFound = errors.New("Group not found")
)

// GroupService handles communication with the group related
// methods of the STNS API.
//
// STNS API docs: http://stns.jp/
type GroupService struct {
	client *Client
}

// List the all groups in STNS Server.
//
// STNS API docs: Now not exist
func (u *GroupService) List() (stns.Attributes, error) {
	req, err := u.client.NewRequest("GET", "group/list", nil)
	if err != nil {
		return nil, err
	}

	response := new(stns.ResponseFormat)
	err = u.client.Do(req, &response)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

// Get group with specified name in STNS Server.
//
// STNS API docs: Now not exist
func (u *GroupService) Get(name string) (*stns.Attribute, error) {
	req, err := u.client.NewRequest("GET", fmt.Sprintf("group/name/%v", name), nil)
	if err != nil {
		return nil, err
	}

	response := new(stns.ResponseFormat)
	err = u.client.Do(req, &response)
	if err != nil {
		return nil, err
	}

	for k, v := range response.Items {
		if k != name {
			continue
		}
		return v, nil
	}

	return nil, ErrGroupNotFound
}

// Get group with specified id in STNS Server.
//
// STNS API docs: Now not exist
func (u *GroupService) GetById(id int) (*stns.Attribute, error) {
	req, err := u.client.NewRequest("GET", fmt.Sprintf("group/id/%v", id), nil)
	if err != nil {
		return nil, err
	}

	response := new(stns.ResponseFormat)
	err = u.client.Do(req, &response)
	if err != nil {
		return nil, err
	}

	for _, v := range response.Items {
		return v, nil
	}

	return nil, ErrGroupNotFound
}
