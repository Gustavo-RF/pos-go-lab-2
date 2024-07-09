package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZipCodeResponse(t *testing.T) {
	tests := []struct {
		name       string
		localidade string
		expect     ZipCodeResponse
	}{
		{
			"Valid location",
			"São Paulo",
			ZipCodeResponse{Localidade: "São Paulo"},
		},
		{
			"Empty location",
			"",
			ZipCodeResponse{Localidade: ""},
		},
		{
			"Different location",
			"Rio de Janeiro",
			ZipCodeResponse{Localidade: "Rio de Janeiro"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NewZipCodeResponse(test.localidade)
			assert.Equal(t, test.expect, result)
		})
	}
}
