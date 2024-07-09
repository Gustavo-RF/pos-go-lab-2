package zipcode

import (
	"errors"
	"testing"

	"github.com/Gustavo-RF/pos-go-lab-2/service-b/internal/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetZipCode(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return([]byte(`{"localidade": "São Paulo"}`), nil)

	resp, err := GetZipCode("01001000", mockRequestFunc.Request)
	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", resp.Localidade)

	mockRequestFunc.AssertExpectations(t)
}

func TestGetWeather_RequestError(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return(nil, errors.New("request error"))

	_, err := GetZipCode("01001000", mockRequestFunc.Request)
	assert.Error(t, err)
	assert.Equal(t, "request error", err.Error())

	mockRequestFunc.AssertExpectations(t)
}

func TestFetch(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return([]byte(`{"localidade": "São Paulo"}`), nil)

	resp, err := fetch("01001000", mockRequestFunc.Request)
	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", resp.Localidade)

	mockRequestFunc.AssertExpectations(t)
}

func TestFetch_RequestError(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return(nil, errors.New("request error"))

	_, err := fetch("01001000", mockRequestFunc.Request)
	assert.Error(t, err)
	assert.Equal(t, "request error", err.Error())

	mockRequestFunc.AssertExpectations(t)
}

func TestFetch_ZipCodeNotFound(t *testing.T) {
	mockRequester := new(web.MockRequestFunc)
	mockRequester.On("Request", mock.Anything, "GET").Return([]byte(`{"erro": "true"}`), nil)

	_, err := fetch("00000000", mockRequester.Request)
	assert.Error(t, err)
	assert.Equal(t, "zipcode not found", err.Error())

	mockRequester.AssertExpectations(t)
}
