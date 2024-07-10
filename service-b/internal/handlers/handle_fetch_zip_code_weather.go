package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gustavo-RF/pos-go-lab-2/service-b/internal/web"
	"github.com/Gustavo-RF/pos-go-lab-2/service-b/weather"
	zipcode "github.com/Gustavo-RF/pos-go-lab-2/service-b/zip-code"
)

type Response struct {
	Message string `json:"message"`
}

func HandleFetchZipCodeTemp(res http.ResponseWriter, req *http.Request, weatherApiKey string) {
	cep := req.URL.Query().Get("cep")

	cepFind, err := zipcode.GetZipCode(cep, web.Request)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	weather, err := weather.GetWeather(cepFind.Localidade, web.Request, weatherApiKey)

	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while get weather: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.Header().Set("Content-type", "application/json")
	fmt.Println("Response: %v", weather)
	json.NewEncoder(res).Encode(weather)
}
