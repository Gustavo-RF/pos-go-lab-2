package main

import (
	"fmt"
	"net/http"

	"github.com/Gustavo-RF/pos-go-lab-2/service-b/configs"
	"github.com/Gustavo-RF/pos-go-lab-2/service-b/internal/handlers"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleFetchZipCodeTemp(w, r, configs.WeatherApiKey)
	})

	fmt.Println("Server started at 8081")
	http.ListenAndServe(":8081", nil)
}
