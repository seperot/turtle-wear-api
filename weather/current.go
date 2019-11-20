package weather

import "github.com/seperot/turtle-wear-api.git/getjson"

type CurrentWeather struct {
	Temp string `json:"Temp"`
	Icon string `json:"Icon"`
}

type RetrieverFunc func(lan string, lon string, json getjson.RetrieveJson) (string, string)

func Getter(lat string, lon string, fn RetrieverFunc, json getjson.RetrieveJson) CurrentWeather {
	condition, temp := fn(lat, lon, json)
	return CurrentWeather{temp, condition}
}
