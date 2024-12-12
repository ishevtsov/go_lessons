package main

import "fmt"

type HTTPMethod int

const (
	GET     HTTPMethod = 0
	POST    HTTPMethod = 1
	PUT     HTTPMethod = 2
	DELETE  HTTPMethod = 3
	PATCH   HTTPMethod = 4
	HEAD    HTTPMethod = 5
	OPTIONS HTTPMethod = 6
	TRACE   HTTPMethod = 7
	CONNECT HTTPMethod = 8
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
