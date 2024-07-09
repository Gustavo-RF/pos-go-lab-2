package entities

import (
	"encoding/json"
)

type WeatherApiResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Coutry string `json:"country"`
}

type Current struct {
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
}

func NewWeatherApiResponse(weather []byte) (*WeatherApiResponse, error) {
	var data WeatherApiResponse
	err := json.Unmarshal(weather, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
