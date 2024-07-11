package entities

import (
	"service-b/internal/web"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestShouldReturnValidField(t *testing.T) {
	res, err := web.Request("https://viacep.com.br/ws/29092260/json/", "GET")

	assert.Nil(t, err)

	zipCodeApiResponse, err := NewZipCodeApiResponse(res)

	assert.Nil(t, err)

	assert.Equal(t, zipCodeApiResponse.Cep, "29092-260")
	assert.Equal(t, zipCodeApiResponse.Localidade, "Vitória")
}
