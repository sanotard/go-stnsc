package stnsc

import (
	"errors"
	"fmt"

	"github.com/STNS/STNS/stns"
	"github.com/sona-tar/go-stnsc/crypto"
)

var (
	// User not found
	ErrUserNotFound = errors.New("User not found")
)

// UserService handles communication with the user related
// methods of the STNS API.
//
// STNS API docs: http://stns.jp/
type UserService struct {
	client *Client
}

// List the all users in STNS Server.
//
// STNS API docs: Now not exist
func (u *UserService) List() (stns.Attributes, error) {
	req, err := u.client.NewRequest("GET", "user/list", nil)
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

// Get user with specified name in STNS Server.
//
// STNS API docs: Now not exist
func (u *UserService) Get(name string) (*stns.Attribute, error) {
	req, err := u.client.NewRequest("GET", fmt.Sprintf("user/name/%v", name), nil)
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

	return nil, ErrUserNotFound
}

// Get user with specified id in STNS Server.
//
// STNS API docs: Now not exist
func (u *UserService) GetById(id int) (*stns.Attribute, error) {
	req, err := u.client.NewRequest("GET", fmt.Sprintf("user/id/%v", id), nil)
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

	return nil, ErrUserNotFound
}

// Authenticate user with specified name and passoword in STNS Server.
//
// STNS API docs: Now not exist
func (u *UserService) Authenticate(username string, password string) (*stns.Attribute, error) {
	attr, err := u.Get(username)
	if err != nil {
		return nil, err
	}

	err = crypto.Verify(attr.User.Password, []byte(password))
	if err != nil {
		return nil, err
	}

	return attr, nil
}
