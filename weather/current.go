package weather

type CurrentWeather struct {
	Temp string `json:"Temp"`
	Icon string `json:"Icon"`
}

func Getter (lat string, lon string) CurrentWeather {
	condition, temp := openWeather(lat, lon)
	return CurrentWeather{temp, condition}
}
