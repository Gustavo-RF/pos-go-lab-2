package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/paemuri/brdoc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Request struct {
	Cep string `json:"cep"`
}

type Response struct {
	Message string `json:"message"`
}

type WeatherRequest struct {
	Tr  trace.Tracer `json:"tr"`
	Cep string       `json:"cep"`
}

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

func Handler(w http.ResponseWriter, r *http.Request, t trace.Tracer, ctx context.Context) {

	ctx, span := t.Start(ctx, "check-cep")
	defer span.End()

	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while fetch zip code data",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if isValid := validate(w, request.Cep); !isValid {
		return
	}

	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	out, err := json.Marshal(WeatherRequest{
		Tr:  t,
		Cep: request.Cep,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while fetch zip code data",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "http://host.docker.internal:8081", bytes.NewBuffer(out))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accepts", "application/json")

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data WeatherResponse
	err = json.Unmarshal(res, &data)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(data)
}

func validate(res http.ResponseWriter, cep string) bool {
	if cep == "" {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Cep is required",
		}
		json.NewEncoder(res).Encode(response)
		return false
	}

	if len(cep) != 8 {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Invalid zipcode",
		}
		json.NewEncoder(res).Encode(response)
		return false
	}

	if !brdoc.IsCEP(cep) {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Invalid zipcode",
		}
		json.NewEncoder(res).Encode(response)
		return false
	}

	return true
}
