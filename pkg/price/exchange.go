package price

import (
	"fmt"
	"github.com/seperot/turtle-wear-api.git/pkg/getjson"
	"net/http"
	"time"
)

func newClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 2,
	}
}

func TradeOgre(json getjson.RetrieveJson) string {
	url := "https://tradeogre.com/api/v1/ticker/BTC-TRTL"
	return fmt.Sprintf("%v", json(url, newClient())["price"])
}

func kuCoin(json getjson.RetrieveJson) string {
	url := "https://openapi-v2.kucoin.com/api/v1/market/stats?symbol=TRTL-BTC"
	return fmt.Sprintf("%v", json(url, newClient())["data"].(map[string]interface{})["last"])
}

func BtcFiatPrice(currency string, json getjson.RetrieveJson) string {
	url := "https://blockchain.info/ticker"
	if currency == "BTC" {
		return "1"
	}
	return fmt.Sprintf("%v", json(url, newClient())[currency].(map[string]interface{})["last"])
}