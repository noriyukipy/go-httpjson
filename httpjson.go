package httpjson

import (
	"encoding/json"
	"net/http"
	"strings"
)

type validationError struct{}

// Validate checks HTTP request method and header.
// Validate returns both error and corresponded StatusCode when validation fails.
// The StatusCode can be set as a status code for a HTTP response.
func Validate(req *http.Request) (error, int) {
	// Check request method
	if req.Method != http.MethodPost {
		// return 405 Method Not Allowed
		return &validationError{}, http.StatusMethodNotAllowed
	}

	// Check Content-Type
	contentTypeOk := false
	for _, contentType := range strings.Split(req.Header.Get("Content-Type"), ";") {
		contentTypeTrimed := strings.TrimSpace(contentType)
		if contentTypeTrimed == "application/json" {
			contentTypeOk = true
		}
	}
	if !contentTypeOk {
		// return 405 Bad Request
		return &validationError{}, http.StatusBadRequest
	}

	return nil, http.StatusOK
}

func Decode(req *http.Request, v interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}

func (err *validationError) Error() string {
	return "ValidationError"
}

func ValidateAndDecode(req *http.Request, v interface{}) (error, int) {
	if err, statusCode := Validate(req); err != nil {
		return &validationError{}, statusCode
	}
	if err := Decode(req, &v); err != nil {
		return err, http.StatusBadRequest
	}

	return nil, http.StatusOK
}
