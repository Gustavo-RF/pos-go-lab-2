package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/paemuri/brdoc"
)

type Request struct {
	Cep string `json:"cep"`
}

type Response struct {
	Message string `json:"message"`
}

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

		validate(w, request.Cep)

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:8081?cep=%s", request.Cep), nil)

		if err != nil {
			panic(err)
		}

		req.Header.Set("Accepts", "application/json")

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
	})

	fmt.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}

func validate(res http.ResponseWriter, cep string) {
	if cep == "" {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Cep is required",
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	if len(cep) != 8 {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Invalid zipcode",
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	if !brdoc.IsCEP(cep) {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Invalid zipcode",
		}
		json.NewEncoder(res).Encode(response)
		return
	}
}
