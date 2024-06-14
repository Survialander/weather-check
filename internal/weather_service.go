package internal

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Survialander/weather/internal/utils"
	"github.com/spf13/viper"
)

type WeatherService struct{}

type WeatherServiceInterface interface {
	GetWeatherData(city string) (WeatherData, error)
}

type WeatherData struct {
	Current struct {
		Temp_c float32
		Temp_f float32
	}
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetWeatherData(city string) (WeatherData, error) {
	apiKey := viper.GetString("WEATHER_KEY")
	url, _ := url.Parse("https://api.weatherapi.com/v1/current.json")

	qParams := url.Query()
	qParams.Add("q", city)
	qParams.Add("key", apiKey)
	url.RawQuery = qParams.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return WeatherData{}, err
	}

	client := utils.GetHttpClient()

	response, err := client.Do(req)
	if err != nil {
		return WeatherData{}, err
	}

	var data WeatherData
	err = json.NewDecoder(response.Body).Decode(&data)

	return data, err
}
