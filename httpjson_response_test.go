package httpjson

import (
	"net/http"
	"testing"
)

// Mock for ResponseWriter
type ResponseWriterMock struct {
	statusCode int
	header http.Header
}

func (self *ResponseWriterMock) Header() http.Header {
	return self.header
}

func (self *ResponseWriterMock) Write(v []byte) (int, error) {
	return 0, nil
}

func (self *ResponseWriterMock) WriteHeader(statusCode int) {
	self.statusCode = statusCode
}

func newResponseWriterMock() *ResponseWriterMock {
	return &ResponseWriterMock{statusCode: http.StatusOK, header: make(http.Header)}
}

// newTestStruct is struct for testing ResposneWriter
type testStruct struct {
	Token string `json:"token"`
	App   string `json:"app"`

}

func newTestStruct() *testStruct {
	return &testStruct{Token:"TOKEN", App: "APP"}
}

// Test for ResponseWriter
// Case: Valid params
func TestWriteResponse(t *testing.T) {
	rw := newResponseWriterMock()
	v := newTestStruct()
	err := WriteResponse(rw, http.StatusOK, v)
	if err != nil {
		t.Errorf("WriteResponse error %s", err)
	}
}