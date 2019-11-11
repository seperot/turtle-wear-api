package main

import (
	"encoding/json"
	"github.com/seperot/turtle-wear-api.git/price"
	"log"
	"net/http"
)

func serveCoinValue(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(price.PriceCalc())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ , err = w.Write(js)
	default:
		w.WriteHeader(http.StatusNotFound)
		_ , err = w.Write([]byte(`{"error": "nothing found"}`))
	}
}

func main() {
	http.HandleFunc("/coin", serveCoinValue)
	http.HandleFunc("/weather", serveCoinValue)
	log.Fatal(http.ListenAndServe(":3000", nil))
}