package main

import (
	"encoding/json"
	"log"
	"net/http"
	"trtlFunctions/priceCalculator"
)

type coinValue struct {
	Usd string `json:"USD"`
	Btc string `json:"BTC"`
}

func serveCoinValue(w http.ResponseWriter, r *http.Request) {
	coinval := coinValue{"$" + priceCalculator.PriceCalc("USD", priceCalculator.ExchangeOgre),
		"₿" + priceCalculator.PriceCalc("BTC", priceCalculator.ExchangeOgre)}

	js, err := json.Marshal(coinval)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/", serveCoinValue)
	log.Fatal(http.ListenAndServe(":8080", nil))
}