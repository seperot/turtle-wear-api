package weather

import (
	"github.com/seperot/turtle-wear-api.git/pkg/getjson"
	"math"
	"strconv"
)

type CurrentWeather struct {
	Temp string `json:"Temp"`
	Icon string `json:"Icon"`
}

type RetrieverFunc func(lan string, lon string, json getjson.RetrieveJson) (string, string)

func Getter(lat string, lon string, fn RetrieverFunc, json getjson.RetrieveJson) CurrentWeather {
	condition, temp := fn(lat, lon, json)
	round, _ := strconv.ParseFloat(temp, 64)
	strconv.FormatFloat(math.Round(round),'f',0, 64)
	return CurrentWeather{strconv.FormatFloat(math.Round(round),'f',0, 64), condition}
}
