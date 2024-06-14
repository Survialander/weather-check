package main

import (
	"net/http"

	"github.com/Survialander/weather/configs"
	"github.com/Survialander/weather/internal"
	"github.com/Survialander/weather/internal/handlers"
	"github.com/go-chi/chi"
)

func main() {
	err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	cepService := internal.NewCepService()
	weatherService := internal.NewWeatherService()
	weatherHandler := handlers.NewWeatherHandler(cepService, weatherService)

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		weatherHandler.GetWeather(w, r)
	})

	http.ListenAndServe(":8080", router)
}
