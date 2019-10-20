package priceCalculator

import "trtlFunctions/utlis"

func tradeOgre() string {
	url := "https://tradeogre.com/api/v1/ticker/BTC-TRTL"
	result := utlis.GetJsonMap(url)
	return result["price"].(string)
}

func kuCoin() string {
	url := "https://openapi-v2.kucoin.com/api/v1/market/stats?symbol=TRTL-BTC"
	result := utlis.GetJsonMap(url)["data"].(map[string]interface{})
	return result["last"].(string)
}

func btcFiatPrice(currency string) float64 {
	url := "https://blockchain.info/ticker"
	if currency == "BTC" {
		return 1.0
	}
	result := utlis.GetJsonMap(url)[currency].(map[string]interface{})
	return result["last"].(float64)
}