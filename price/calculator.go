package price

import (
	"strconv"
)

type CoinValue struct {
	Usd string `json:"USD"`
	Btc string `json:"BTC"`
}

type CryptoFunc func() string
type FiatFunc func(fiat string) string

func Calc(cryptoFn CryptoFunc, fiatFn FiatFunc) CoinValue {
	fiatPrice, _ := strconv.ParseFloat(fiatFn("USD"), 64)
	exchangePrice, _ := strconv.ParseFloat(cryptoFn(), 64)
	return CoinValue{"$" + strconv.FormatFloat(fiatPrice * exchangePrice, 'f', 9, 64),
		"₿" + strconv.FormatFloat(exchangePrice,'f',9, 64)}
}