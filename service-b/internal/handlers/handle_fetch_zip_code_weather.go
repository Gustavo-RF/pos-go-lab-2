package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service-b/internal/web"
	"service-b/weather"
	zipcode "service-b/zip-code"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Request struct {
	Tr  trace.Tracer `json:"tr"`
	Cep string       `json:"cep"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleFetchZipCodeTemp(res http.ResponseWriter, req *http.Request, t trace.Tracer, weatherApiKey string) {

	carrier := propagation.HeaderCarrier(req.Header)
	ctx := req.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)

	// tr := otel.GetTracerProvider().Tracer("component-main-2")
	ctx, span := t.Start(ctx, "check-cep-2")
	defer span.End()

	if err != nil {
		fmt.Printf("Chegou aqui err not nil: %s", err.Error())
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while fetch data: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	cepFind, err := zipcode.GetZipCodeWithContext(ctx, request.Cep, web.RequestWithContext)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	weather, err := weather.GetWeatherWithContext(ctx, cepFind.Localidade, web.RequestWithContext, weatherApiKey)

	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while get weather: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.Header().Set("Content-type", "application/json")

	json.NewEncoder(res).Encode(weather)
}
