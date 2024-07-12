package zipcode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"service-b/zip-code/entities"

	"go.opentelemetry.io/otel/trace"
)

type RequestFunc func(url, method string) ([]byte, error)
type RequestWithContextFunc func(ctx context.Context, url, method string) ([]byte, error)

func GetZipCode(zipcode string, requestFunc RequestFunc) (*entities.ZipCodeResponse, error) {

	zipCodeApiResponse, err := fetch(zipcode, requestFunc)

	if err != nil {
		return nil, err
	}

	response := entities.NewZipCodeResponse(zipCodeApiResponse.Localidade)

	return &response, nil
}

func GetZipCodeWithContext(ctx context.Context, zipcode string, requestFunc RequestWithContextFunc, tracer trace.Tracer, req *http.Request) (*entities.ZipCodeResponse, error) {

	ctx, span := tracer.Start(req.Context(), "Service B - Zip code - Start Tracer")
	defer span.End()

	zipCodeApiResponse, err := fetchWitchContext(ctx, zipcode, requestFunc)

	if err != nil {
		return nil, err
	}

	response := entities.NewZipCodeResponse(zipCodeApiResponse.Localidade)

	return &response, nil
}

func fetch(zipcode string, requestFunc RequestFunc) (*entities.ZipCodeApiResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	res, err := requestFunc(url, "GET")

	if err != nil {
		return nil, err
	}

	zipCodeApiResponse, err := entities.NewZipCodeApiResponse(res)
	if err != nil {
		return nil, err
	}

	if zipCodeApiResponse.Erro == "true" {
		return nil, errors.New("zipcode not found")
	}

	return zipCodeApiResponse, nil
}

func fetchWitchContext(ctx context.Context, zipcode string, requestFunc RequestWithContextFunc) (*entities.ZipCodeApiResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	res, err := requestFunc(ctx, url, "GET")

	if err != nil {
		return nil, err
	}

	zipCodeApiResponse, err := entities.NewZipCodeApiResponse(res)
	if err != nil {
		return nil, err
	}

	if zipCodeApiResponse.Erro == "true" {
		return nil, errors.New("zipcode not found")
	}

	return zipCodeApiResponse, nil
}
