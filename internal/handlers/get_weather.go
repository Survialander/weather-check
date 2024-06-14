package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Survialander/weather/internal"
)

type WeatherHandler struct {
	CepService     internal.CepServiceInterface
	WeatherService internal.WeatherServiceInterface
}

type WeatherResponse struct {
	Temp_c float32 `json:"temp_c"`
	Temp_f float32 `json:"temp_f"`
	Temp_k float32 `json:"temp_k"`
}

func NewWeatherHandler(cepService internal.CepServiceInterface, weatherService internal.WeatherServiceInterface) *WeatherHandler {
	return &WeatherHandler{
		CepService:     cepService,
		WeatherService: weatherService,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	cepData, err := h.CepService.GetCepData(cep)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if cepData.Localidade == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	weatherData, err := h.WeatherService.GetWeatherData(cepData.Localidade)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := WeatherResponse{
		Temp_c: weatherData.Current.Temp_c,
		Temp_f: weatherData.Current.Temp_f,
		Temp_k: weatherData.Current.Temp_c + 273,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
