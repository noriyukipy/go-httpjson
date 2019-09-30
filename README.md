# go-httpjson

[![GoDoc](https://godoc.org/github.com/noriyukipy/go-httpjson?status.svg)](https://godoc.org/github.com/noriyukipy/go-httpjson)

This repository provides useful functions to deal with JSON requests via HTTP.

## Getting Started

### ValidateAndDecode

`ValidateAndDecode` function checks HTTP request with JSON body and decode JSON to struct.

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
