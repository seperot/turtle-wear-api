package price

import "testing"

func TestCalc(t *testing.T) {
	result := Calc(testCrypto, testFiat)
	if result.Usd != "$0.030703290"{
		t.Error("got", result)
	}
}

func testCrypto() string {
return "0.000003"
}

func testFiat(currency string) string {
return "10234.43"
}