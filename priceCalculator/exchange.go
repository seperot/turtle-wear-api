package priceCalculator

import (
	"fmt"
	"github.com/seperot/turtle-wear-api.git/utlis"
)

func tradeOgre() string {
	url := "https://tradeogre.com/api/v1/ticker/BTC-TRTL"
	return fmt.Sprintf("%v", utlis.GetJsonMap(url)["price"])
}

func kuCoin() string {
	url := "https://openapi-v2.kucoin.com/api/v1/market/stats?symbol=TRTL-BTC"
	return fmt.Sprintf("%v", utlis.GetJsonMap(url)["data"].(map[string]interface{})["last"])
}

func btcFiatPrice(currency string) string {
	url := "https://blockchain.info/ticker"
	if currency == "BTC" {
		return "1"
	}
	return fmt.Sprintf("%v", utlis.GetJsonMap(url)[currency].(map[string]interface{})["last"])
}