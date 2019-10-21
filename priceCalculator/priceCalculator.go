package priceCalculator

import (
	"strconv"
)
const ExchangeOgre int = 1
const ExchangeKu int = 2

func PriceCalc(currency string, exchange int) string {
	exchangePrice, _ := strconv.ParseFloat(parseExchange(exchange), 64)
	fiatPrice, _ := strconv.ParseFloat(btcFiatPrice(currency), 64)
	return strconv.FormatFloat(fiatPrice * exchangePrice, 'f', 9, 64)
}

func parseExchange(exchange int) string {
	switch exchange {
	case 1:
		return tradeOgre()
	case 2:
		return kuCoin()
	default:
		return "0"
	}
}