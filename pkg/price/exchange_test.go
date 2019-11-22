package price

import (
	"encoding/json"
	"log"
	"testing"
)

var stubbedResponse []byte
func TestBtcFiatPrice(t *testing.T) {
	stubbedResponse = []byte(`{"USD":{"last": "12.1"}}`)
	result1 := BtcFiatPrice("BTC", testJsonRetriever)
	result2 := BtcFiatPrice("USD", testJsonRetriever)
	if result1 != "1" || result2 != "12.1" {
		t.Error("OH NO")
	}
}

func TestTradeOgre(t *testing.T) {
	stubbedResponse = []byte(`{"price": "0.00001"}`)
	result := TradeOgre(testJsonRetriever)
	if result != "0.00001" {
		t.Error("OH NO")
	}
}

func Test_kuCoin(t *testing.T) {
	stubbedResponse = []byte(`{"data": {"last": "12.1"}}`)
	result := kuCoin(testJsonRetriever)
	if result != "12.1" {
		t.Error("OH NO")
	}
}

func testJsonRetriever(fullUrl string) map[string]interface{} {
	var result map[string]interface{}
	jsonErr := json.Unmarshal(stubbedResponse, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return result
}