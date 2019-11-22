package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleEventUnknownRequest(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/error", nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventUnknownMethod(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/coin", nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventValidRequest(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/coin", nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventValidRequestNoHeader(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/weather", nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventValidRequestWithHeader(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/weather", nil)
	req.Header.Set("lat", "12")
	req.Header.Set("lon", "12")
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventUnknownMethodWithHeader(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/weather", nil)
	req.Header.Set("lat", "12")
	req.Header.Set("lon", "12")
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}