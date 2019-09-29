# go-httpjson

This repository provides useful functions to deal with JSON requests via HTTP.

## Overview

`ValidateAndDecode` function checks HTTP request with JSON body and decode JSON to struct.

```go
v := &struct {
	Token string `json:"token"`
	App   string `json:"app"`
}{}
err, statusCode := httpjson.ValidateAndDecode(req, &v)
```
