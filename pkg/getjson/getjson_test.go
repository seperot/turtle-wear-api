package getjson

import (
	"fmt"
	"net/http"
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
	go func() {
		_ = http.ListenAndServe(":8080", handler)
	}()
	result := fmt.Sprintf("%v", Map("http://:8080")["Test"].(map[string]interface{})["Jeff"])
	if result != "Heff" {
		t.Error("OH NO")
	}
}