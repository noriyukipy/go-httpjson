package httpjson

import (
	"bytes"
	"net/http"
	"testing"
)

// HTTP request factory functions
func buildValidRequest() *http.Request {
	jsonBody := `{"token": "TOKEN", "app": "APP"}`
	req, _ := http.NewRequest(http.MethodPost, "http://example.com", bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestValidateOK(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "http://example.com", nil)
	req.Header.Set("Content-Type", "application/json")

	err, code := Validate(req)
	if !(err == nil && code == http.StatusOK) {
		t.Error("Valid request validation failed")
	}
}

func TestValidateInvalidMethod(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	err, code := Validate(req)
	if !(err != nil && code == http.StatusMethodNotAllowed) {
		t.Error("Method validation failed")
	}
}

func TestValidateInvalidContentType(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	req.Header.Set("Content-Type", "text/plain")

	err, code := Validate(req)
	if !(err != nil && code == http.StatusMethodNotAllowed) {
		t.Error("Content-Type validation failed")
	}
}

func TestDecodeJSONValidBody(t *testing.T) {
	req := buildValidRequest()

	v := &struct {
		Token string `json:"token"`
		App   string `json:"app"`
	}{}
	err := Decode(req, &v)
	if err != nil {
		t.Error("Valid JSON decoding fail ")
	}
	if !(v.Token == "TOKEN" && v.App == "APP") {
		t.Error("Valid JSON decoding fail ")
	}
}

func TestDecodeJSONInvalidBody(t *testing.T) {
	// Build informal JSON request
	jsonBody := `{"token": "TOKEN"`
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Set("Content-Type", "text/plain")

	v := &struct {
		Token string `json:"token"`
		App   string `json:"app"`
	}{}
	err := Decode(req, &v)
	if err == nil {
		t.Error("Invalid JSON decoding fail ")
	}
}

func TestValidateAndDecode(t *testing.T) {
	req := buildValidRequest()
	v := &struct {
		Token string `json:"token"`
		App   string `json:"app"`
	}{}
	err, _ := ValidateAndDecode(req, &v)
	t.Log(err)
	if err != nil {
		t.Error("ValidateAndDecode fail for valid request")

	}
}
