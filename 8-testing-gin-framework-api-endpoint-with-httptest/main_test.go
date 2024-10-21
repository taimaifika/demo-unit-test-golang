package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinPingEndpoint(t *testing.T) {
	// router
	router := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// check status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	// check response body
	if w.Body.String() != "pong" {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), "pong")
	}
}
