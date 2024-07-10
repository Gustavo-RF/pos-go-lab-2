package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gustavo-RF/pos-go-lab-2/service-b/internal/web"
	"github.com/Gustavo-RF/pos-go-lab-2/service-b/weather"
	zipcode "github.com/Gustavo-RF/pos-go-lab-2/service-b/zip-code"
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

func HandleFetchZipCodeTemp(res http.ResponseWriter, req *http.Request, weatherApiKey string) {

	var request Request
	err := json.NewDecoder(req.Body).Decode(&request)
	fmt.Printf("Chegou aqui 111: %v", req.Body)

	if err != nil {
		fmt.Printf("Chegou aqui err not nil: %s", err.Error())
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while fetch data: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	fmt.Printf("Chegou aqui 1:")
	carrier := propagation.HeaderCarrier(req.Header)
	ctx := req.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	fmt.Printf("Chegou aqui:")
	ctx, span := request.Tr.Start(ctx, "check-cep-2", trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

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
