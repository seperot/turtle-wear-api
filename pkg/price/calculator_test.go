package price

import (
	"github.com/seperot/turtle-wear-api.git/pkg/getjson"
	"testing"
)

func TestCalc(t *testing.T) {
	result := Calc(testCrypto, testFiat, getjson.Map)
	if result.Usd != "$0.030703290"{
		t.Error("got", result)
	}
}

func testCrypto(json getjson.RetrieveJson) string {
return "0.000003"
}

func testFiat(currency string, json getjson.RetrieveJson) string {
return "10234.43"
}