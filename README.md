# go-httpjson

[![GoDoc](https://godoc.org/github.com/noriyukipy/go-httpjson?status.svg)](https://godoc.org/github.com/noriyukipy/go-httpjson)

This repository provides useful functions to deal with JSON requests via HTTP.

## Getting Started

### ValidateAndDecode

`ValidateAndDecode` first checks HTTP request method and header, then decode JSON body to struct.
`ValidateAndDecode` returns both error and corresponded StatusCode.

When validation or decode fails, the StatusCode has corresponded error code.
Therefore, the StatusCode can be set as a status code for a HTTP response.

StatusCode is set to `http.StatusOK` if the both of them are success.

```go
type User struct {
	Name string `json:"text"`
	Age int `json:"int"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err, statusCode := httpjson.ValidateAndDecode(r, &user)
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(statusCode)
		return
	}
	log.Printf("Request: %v", user)
}
```
