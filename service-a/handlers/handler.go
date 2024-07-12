package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/paemuri/brdoc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Request struct {
	Cep string `json:"cep"`
}

type Response struct {
	Message string `json:"message"`
}

type WeatherRequest struct {
	Cep string `json:"cep"`
}

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

const name = "Service A - Tracer"

var (
	tracer = otel.Tracer(name)
)

func Handler(w http.ResponseWriter, r *http.Request) {

	ctx, span := tracer.Start(r.Context(), "Start Tracer")
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

	ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	out, err := json.Marshal(WeatherRequest{
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

	time.Sleep(time.Second * 1)
	fmt.Println("Call service b..")
	req, err := http.NewRequestWithContext(ctx, "POST", "http://host.docker.internal:8081", bytes.NewBuffer(out))

	if err != nil {
		return
	}

	req.Header.Set("Accepts", "application/json")

	req.Close = true

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	fmt.Println("Called service b..")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data WeatherResponse
	err = json.Unmarshal(res, &data)

	if err != nil {
		return
	}

	time.Sleep(time.Second * 1)
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
