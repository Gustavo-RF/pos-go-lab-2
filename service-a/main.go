package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paemuri/brdoc"
)

type Request struct {
	Cep string `json:"cep"`
}

type Response struct {
	Message string `json:"message"`
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
