package main

import (
	"encoding/json"
	"github.com/seperot/turtle-wear-api.git/pkg/getjson"
	"github.com/seperot/turtle-wear-api.git/pkg/price"
	"github.com/seperot/turtle-wear-api.git/pkg/weather"
	"net/http"
)

func handleEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			js []byte
			err error
		)

		switch r.URL.Path {
		case "/coin":
			js, err = json.Marshal(price.Calc(price.TradeOgre, price.BtcFiatPrice, getjson.Map))
		case "/weather":
			lat := r.Header.Get("lat")
			lon := r.Header.Get("lon")
			if len(lat) <= 0 || len(lon) <= 0{
				http.Error(w, "Missing Lat Lon Headers", http.StatusInternalServerError)
			} else {
				js, err = json.Marshal(weather.Getter(lat, lon, weather.OpenWeather, getjson.Map))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			_ , err = w.Write([]byte(`{"error": "nothing found"}`))
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

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	err := http.ListenAndServe(":3000", handleEvent())
	if err != nil {
		panic(err)
	}
}