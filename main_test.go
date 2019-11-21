package main

import (
	"net/http"
	"testing"
)

func Test_handleEvent(t *testing.T) {
	go func() {
		_ = http.ListenAndServe(":9021", handleEvent())
	}()
	req, _ := http.NewRequest("GET", "http://localhost:9021/noterror", nil)
	if req != nil {

	}
}