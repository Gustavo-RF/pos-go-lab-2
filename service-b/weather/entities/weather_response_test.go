package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name   string
		city   string
		tempC  float32
		expect float32
	}{
		{"Freezing point", "Vitória", 0, 32},
		{"Boiling point", "Vitória", 100, 212},
		{"Negative temperature", "Vitória", -40, -40},
		{"Room temperature", "Vitória", 25, 77},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ConvertCelsiusToFahrenheit(test.tempC)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name   string
		city   string
		tempC  float32
		expect float32
	}{
		{"Freezing point", "Vitória", 0, 273},
		{"Boiling point", "Vitória", 100, 373},
		{"Negative temperature", "Vitória", -273, 0},
		{"Room temperature", "Vitória", 25, 298},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertCelsiusToKelvin(tt.tempC)
			assert.Equal(t, tt.expect, result)
		})
	}
}

func TestNewWeatherResponse(t *testing.T) {
	tests := []struct {
		name   string
		city   string
		tempC  float32
		expect WeatherResponse
	}{
		{
			"Freezing point",
			"Vitória",
			0,
			WeatherResponse{City: "Vitória", TempC: 0, TempF: 32, TempK: 273},
		},
		{
			"Boiling point",
			"Vitória",
			100,
			WeatherResponse{City: "Vitória", TempC: 100, TempF: 212, TempK: 373},
		},
		{
			"Negative temperature",
			"Vitória",
			-40,
			WeatherResponse{City: "Vitória", TempC: -40, TempF: -40, TempK: 233},
		},
		{
			"Room temperature",
			"Vitória",
			25,
			WeatherResponse{City: "Vitória", TempC: 25, TempF: 77, TempK: 298},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewWeatherResponse(tt.city, tt.tempC)
			assert.Equal(t, tt.expect, result)
		})
	}
}
