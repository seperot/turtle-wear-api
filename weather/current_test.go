package weather

import (
	"github.com/seperot/turtle-wear-api.git/getjson"
	"testing"
)

func TestGetter(t *testing.T) {
	result := Getter("1","1", testRetriever, getjson.Map)
	if result.Temp != "hot af"{
		t.Error("got", result)
	}
}

func testRetriever(lat string, lon string, json getjson.RetrieveJson) (string, string) {
	return "sunny", "hot af"
}