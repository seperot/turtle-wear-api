package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	err = "/error"
	weth = "/weather"
	crypto = "/coin"
	get = "GET"
	post = "POST"
)

func TestHandleEventUnknownRequest(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(get, err, nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventUnknownMethod(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(post, crypto, nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventValidRequest(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(get, crypto, nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventValidRequestNoHeader(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(get, weth, nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}

func TestHandleEventValidRequestWithHeader(t *testing.T) {
	handler := handleEvent()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(get, weth, nil)
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
	req, err := http.NewRequest(post, weth, nil)
	req.Header.Set("lat", "12")
	req.Header.Set("lon", "12")
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Error("oh no")
	}
}