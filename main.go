package main

import (
	"encoding/json"
	"github.com/seperot/turtle-wear-api.git/getjson"
	"github.com/seperot/turtle-wear-api.git/price"
	"github.com/seperot/turtle-wear-api.git/weather"
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
				return
			}
			js, err = json.Marshal(weather.Getter(lat, lon, weather.OpenWeather, getjson.Map))
		default:
			http.Error(w, "Unknown Response Error", http.StatusInternalServerError)
			return
		}

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
}

func main() {
	panic(http.ListenAndServe(":3000", handleEvent()))
}