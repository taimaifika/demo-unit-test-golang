package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	handler := setupRouter()

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusOK)
	}

	// Check the response body
	if w.Body.String() != "pong" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), "pong")
	}
}
