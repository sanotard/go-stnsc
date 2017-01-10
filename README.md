`go-stnsc` is a SNTS API Client and Authentication Module.
===============================

[![GoDoc](https://godoc.org/github.com/sonatard/go-stnsc?status.svg)](https://godoc.org/github.com/sonatard/go-stnsc/stnsc)

`go-stnsc`  support STNS v2 JSON format.

[SNTS Official page](http://stns.jp/)


## Example

### API

- Client Code:

Server configulation. [stnsc.conf](https://github.com/sonatard/docker-compose-stns/blob/master/stns.conf)
```go

func main() {
	client, err := stnsc.NewClient("http://localhost:1104/v2/", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
```

- Output:
```text
Attributes users : stns.Attributes{"bar":(*stns.Attribute)(0xc4200e01e0), "foo":(*stns.Attribute)(0xc4200e02a0)}
Attribute user : &stns.Attribute{Id:1001, User:(*stns.User)(0xc4200f0180), Group:(*stns.Group)(nil)}
Attribute user : &stns.Attribute{Id:1001, User:(*stns.User)(0xc420152000), Group:(*stns.Group)(nil)}
user : &stns.User{Password:"$6$RNqhn2ttIfMcRj4r$Ddnbckw1T1xUkguDWvSsb3GZseoeahRbr27vKbYV9opja2SKWi6y.67YI0yXz8HremKCpJwwFEOqed6Eic9.0.", GroupId:1002, Directory:"/home/foo", Shell:"/bin/bash", Gecos:"description", Keys:[]string{"key"}, LinkUsers:[]string{"linkuser"}}
Attribute user : &stns.Attribute{Id:1002, User:(*stns.User)(0xc4200f0480), Group:(*stns.Group)(nil)}
user : &stns.User{Password:"$6$gu42K/pg0o7NBP9O$NshQ3iHO4gE3av9.tkE6DWCgA0h1vG1TzH.SHfQn.TEZpmFBVSD0G7pnH3SGKj22RFz5qiy3ezMg6UQ6JJejE.", GroupId:1002, Directory:"/home/bar", Shell:"/bin/bash", Gecos:"description", Keys:[]string{"key"}, LinkUsers:[]string{"linkuser"}}
```


See other API usage. [./example/main.go](./example/main.go)

### Authenticate

You can provide SNTS User Login on your web page easily.

- Client Code:
```go
func main() {
	client, err := stnsc.NewClient("http://localhost:1104/v2/", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
```

- Output:
```text
user : &stns.User{Password:"$6$72qH5tfJta43J1lH$o1OvvIxkDCNZtrAh3UWM9dKkGawTuBeGpLoxRuICH6B/9.Y5PA/bD
tvm.fK/bB8zFNNofus6jQHXzMyiqCCqj0", GroupId:1001, Directory:"/home/example", Shell:"/bin/bash", Gecsoo:"", Keys:[]string{"ssh-rsa XXXXX…"}, LinkUsers:[]string{"foo"}}
```

## Demo
- Create STNS Server
 - See [docker-compose-stns](https://github.com/sonatard/docker-compose-stns)

- Client
```shell
$ go get -v github.com/sonatard/go-stnsc
$ cd ${GOPATH}/src/github.com/sonatard/go-stnsc/example
$ go run main.go
```
