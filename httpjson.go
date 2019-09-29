package httpjson

import (
	"encoding/json"
	"net/http"
)

type StatusCode int
type validationError struct{}

func Validate(req *http.Request) (error, StatusCode) {
	// Check request method
	if req.Method != http.MethodPost {
		// return 405 Method Not Allowed
		return &validationError{}, http.StatusMethodNotAllowed
	}
	// Check content-type
	if req.Header.Get("Content-Type") != "application/json" {
		// return 403
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

func ValidateAndDecode(req *http.Request, v interface{}) (error, StatusCode) {
	if err, statusCode := Validate(req); err != nil {
		return &validationError{}, statusCode
	}
	if err := Decode(req, &v); err != nil {
		return err, http.StatusBadRequest
	}

	return nil, http.StatusOK
}
