// go-stns example package
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/STNS/STNS/stns"
	"github.com/sonatard/go-stnsc/stnsc"
)

func main() {
	url := "http://localhost:1104/v2/"
	// No Auth
	// client, err := stnsc.NewClient(url, nil)

	// Basic Auth
	tp := &stnsc.BasicAuthTransport{
		Username: strings.TrimSpace("basicuser"),
		Password: strings.TrimSpace("basicpass"),
	}

	client, err := stnsc.NewClient(url, tp.Client())
	if err != nil {
		panic(err)
	}

	userAPIExample(client)
	groupAPIExample(client)
	userAuthenticateExample(client)
}

func userAPIExample(client *stnsc.Client) {
	var err error

	var attrUsers stns.Attributes
	attrUsers, err = client.User.List()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attributes users : %#v\n", attrUsers)
	fmt.Printf("Attribute user : %#v\n", attrUsers["foo"])

	var attrUser *stns.Attribute
	attrUser, err = client.User.Get("foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attribute user : %#v\n", attrUser)
	fmt.Printf("user : %#v\n", attrUser.User)

	attrUser, err = client.User.GetById(1002)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attribute user : %#v\n", attrUser)
	fmt.Printf("user : %#v\n", attrUser.User)
}

func groupAPIExample(client *stnsc.Client) {
	var err error

	var attrGroups stns.Attributes
	attrGroups, err = client.Group.List()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attributes groups : %#v\n", attrGroups)
	fmt.Printf("Attribute group : %#v\n", attrGroups["hoge"])

	var attrGroup *stns.Attribute
	attrGroup, err = client.Group.Get("hoge")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attribute group : %#v\n", attrGroup)
	fmt.Printf("group : %#v\n", attrGroup.Group)

	attrGroup, err = client.Group.GetById(2002)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attribute group : %#v\n", attrGroup)
	fmt.Printf("group : %#v\n", attrGroup.Group)

}

func userAuthenticateExample(client *stnsc.Client) {
	var err error

	var attrUser *stns.Attribute
	attrUser, err = client.User.Authenticate("foo", "foopass")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Attribute user : %#v\n", attrUser)
	fmt.Printf("user : %#v\n", attrUser.User)
}
