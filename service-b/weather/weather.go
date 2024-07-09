package weather

import (
	"fmt"
	"net/url"

	"github.com/Gustavo-RF/pos-go-lab-2/service-b/weather/entities"
)

type RequestFunc func(url, method string) ([]byte, error)

func GetWeather(local string, requestFunc RequestFunc, weatherApiKey string) (entities.WeatherResponse, error) {
	weatherApiResponse, err := fetch(local, requestFunc, weatherApiKey)
	if err != nil {
		return entities.WeatherResponse{}, err
	}

	weatherResponse := entities.NewWeatherResponse(weatherApiResponse.Current.TempC)
	return weatherResponse, nil
}

func fetch(local string, requestFunc RequestFunc, weatherApiKey string) (*entities.WeatherApiResponse, error) {
	localEscaped := url.QueryEscape(local)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", weatherApiKey, localEscaped)

	res, err := requestFunc(url, "GET")
	if err != nil {
		return nil, err
	}

	weatherApiResponse, err := entities.NewWeatherApiResponse(res)
	if err != nil {
		return nil, err
	}

	return weatherApiResponse, nil
}
