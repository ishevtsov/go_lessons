package main

import "fmt"

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
	PUT
	DELETE
	PATCH
	HEAD
	OPTIONS
	TRACE
	CONNECT
)

type HTTPRequest struct {
	method  HTTPMethod
	headers map[string]string
	uri     string
}

func main() {
	r := HTTPRequest{method: GET, headers: map[string]string{"Accept": "application/json"}, uri: "/prices"}
	fmt.Println(r)
}
