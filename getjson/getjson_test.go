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
		_ = http.ListenAndServe(":9021", handler)
	}()
	result := fmt.Sprintf("%v", Map("http://localhost:9021")["Test"].(map[string]interface{})["Jeff"])
	if result != "Heff" {
		t.Error("OH NO")
	}
}