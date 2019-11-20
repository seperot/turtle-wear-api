package weather

import (
	"fmt"
	"github.com/seperot/turtle-wear-api.git/getjson"
)

func OpenWeather(lat string, lon string, json getjson.RetrieveJson) (string, string) {
	const (
		url = "https://api.openweathermap.org/data/2.5/weather"
		key = "appid=9a3b7a64c959e1b70a626077efd81b83"
		units = "units=metric"
	)
	latitude := "lat=" + lat
	longitude := "lon=" + lon
	fullUrl := url + "?" + latitude + "&" + longitude + "&" + units + "&" + key
	responseMap := json(fullUrl)
	return fmt.Sprintf("%v", responseMap["weather"].([]interface{})[0].(map[string]interface{})["main"]),
	fmt.Sprintf("%v", responseMap["main"].(map[string]interface{})["temp"])
}
