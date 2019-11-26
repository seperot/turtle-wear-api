package weather

import (
	"github.com/seperot/turtle-wear-api.git/pkg/getjson"
	"testing"
)

func TestGetter(t *testing.T) {
	result := Getter("1","1", testRetriever, getjson.Map)
	if result.Icon != "sunny" || result.Temp != "12"{
		t.Error("got", result)
	}
}

func testRetriever(lat string, lon string, json getjson.RetrieveJson) (string, string) {
	return "sunny", "12.34"
}