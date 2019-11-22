package price

import (
	"fmt"
	"github.com/seperot/turtle-wear-api.git/pkg/getjson"
)

func TradeOgre(json getjson.RetrieveJson) string {
	url := "https://tradeogre.com/api/v1/ticker/BTC-TRTL"
	return fmt.Sprintf("%v", json(url)["price"])
}

func kuCoin(json getjson.RetrieveJson) string {
	url := "https://openapi-v2.kucoin.com/api/v1/market/stats?symbol=TRTL-BTC"
	return fmt.Sprintf("%v", json(url)["data"].(map[string]interface{})["last"])
}

func BtcFiatPrice(currency string, json getjson.RetrieveJson) string {
	url := "https://blockchain.info/ticker"
	if currency == "BTC" {
		return "1"
	}
	return fmt.Sprintf("%v", json(url)[currency].(map[string]interface{})["last"])
}