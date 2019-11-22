package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestOpenWeather(t *testing.T) {
result, result2 := OpenWeather("12", "12", testJsonRetriever)
if result != "test" || result2 != "47" {
	t.Error("OH NO")
}
}

func testJsonRetriever(fullUrl string, client *http.Client) map[string]interface{} {
	stubbedResponse := []byte(`{"weather": [{"main": "test"}, {"main": "fail"}],"main": {"temp": 47}}`)
	var result map[string]interface{}
	jsonErr := json.Unmarshal(stubbedResponse, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return result
}