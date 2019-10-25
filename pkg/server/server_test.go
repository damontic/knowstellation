package server

import (
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/", nil)
	responseWriter := httptest.NewRecorder()
	rootHandler(responseWriter, request)
	if string(responseWriter.Body.Bytes()) != "hello\n" {
		t.Fatalf("Failed")
	}
}
