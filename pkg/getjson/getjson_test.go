package getjson

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler(response []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(response)
	}
}

func TestMapSuccess(t *testing.T) {
	stubbedResponse := []byte(`{"Test": { "Jeff": "Heff"}}`)
	handler := handler(stubbedResponse)
	server := httptest.NewServer(handler)
	result := fmt.Sprintf("%v", Map(server.URL, server.Client())["Test"].(map[string]interface{})["Jeff"])
	if result != "Heff" {
		t.Error("OH NO")
	}
}