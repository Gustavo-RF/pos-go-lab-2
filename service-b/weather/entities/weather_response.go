package entities

type WeatherResponse struct {
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

func NewWeatherResponse(tempc float32) WeatherResponse {
	response := WeatherResponse{
		TempC: tempc,
		TempF: ConvertCelsiusToFahrenheit(tempc),
		TempK: ConvertCelsiusToKelvin(tempc),
	}

	return response
}

func ConvertCelsiusToFahrenheit(tempC float32) float32 {
	return tempC*1.8 + 32
}

func ConvertCelsiusToKelvin(tempC float32) float32 {
	return tempC + 273
}
