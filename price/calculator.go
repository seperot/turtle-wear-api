package price

import (
	"strconv"
)

type CoinValue struct {
	Usd string `json:"USD"`
	Btc string `json:"BTC"`
}

func PriceCalc() CoinValue {
	fiatPrice, _ := strconv.ParseFloat(btcFiatPrice("USD"), 64)
	exchangePrice, exchangePriceErr := strconv.ParseFloat(tradeOgre(), 64)
	if exchangePriceErr != nil {
		exchangeAltPrice, _ := strconv.ParseFloat(kuCoin(), 64)
		return CoinValue{"$" + strconv.FormatFloat(fiatPrice * exchangeAltPrice, 'f', 9, 64),
			"₿" + strconv.FormatFloat(exchangeAltPrice,'f',9, 64)}
	}
	return CoinValue{"$" + strconv.FormatFloat(fiatPrice * exchangePrice, 'f', 9, 64),
		"₿" + strconv.FormatFloat(exchangePrice,'f',9, 64)}
}