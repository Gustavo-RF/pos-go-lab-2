package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service-b/configs"
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

const name = "Service B - Tracer"

var (
	tracer = otel.Tracer(name)
)

func HandleFetchZipCodeTemp(res http.ResponseWriter, req *http.Request) {

	carrier := propagation.HeaderCarrier(req.Header)
	ctx := req.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := tracer.Start(req.Context(), "Service B - Start Tracer")
	defer span.End()

	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while fetch data: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	cepFind, err := zipcode.GetZipCodeWithContext(ctx, request.Cep, web.RequestWithContext, tracer, req)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Call weather\n")
	weather, err := weather.GetWeatherWithContext(ctx, cepFind.Localidade, web.RequestWithContext, configs.WeatherApiKey, tracer, req)
	fmt.Printf("Called weather\n")

	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while get weather: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		fmt.Printf("response: %v", response)
		return
	}

	res.Header().Set("Content-type", "application/json")
	fmt.Println(weather)
	json.NewEncoder(res).Encode(weather)
	fmt.Printf("Before end\n")
	return
}
