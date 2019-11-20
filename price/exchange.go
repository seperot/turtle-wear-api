package price

import (
	"fmt"
	"github.com/seperot/turtle-wear-api.git/getjson"
)

func TradeOgre() string {
	url := "https://tradeogre.com/api/v1/ticker/BTC-TRTL"
	return fmt.Sprintf("%v", getjson.Map(url)["price"])
}

func kuCoin() string {
	url := "https://openapi-v2.kucoin.com/api/v1/market/stats?symbol=TRTL-BTC"
	return fmt.Sprintf("%v", getjson.Map(url)["data"].(map[string]interface{})["last"])
}

func BtcFiatPrice(currency string) string {
	url := "https://blockchain.info/ticker"
	if currency == "BTC" {
		return "1"
	}
	return fmt.Sprintf("%v", getjson.Map(url)[currency].(map[string]interface{})["last"])
}