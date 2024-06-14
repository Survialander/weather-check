package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Survialander/weather/internal"
	"github.com/Survialander/weather/internal/handlers"
	"github.com/Survialander/weather/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldReturnWeatherInformation(t *testing.T) {
	// Arrange
	cepData := internal.CepData{
		Cep:        "00000000",
		Localidade: "Brasilia",
	}
	weatherData := internal.WeatherData{
		Current: struct {
			Temp_c float32
			Temp_f float32
		}{
			Temp_c: 20.8,
			Temp_f: 20.8,
		},
	}
	expectedResponse := handlers.WeatherResponse{
		Temp_c: 20.8,
		Temp_f: 20.8,
		Temp_k: 293.8,
	}

	cepService := mocks.CepServiceMock{}
	weatherService := mocks.WeatherServiceMock{}

	cepService.On("GetCepData", "00000000").Return(cepData, nil)
	weatherService.On("GetWeatherData", "Brasilia").Return(weatherData, nil)
	handler := handlers.NewWeatherHandler(&cepService, &weatherService)

	request := httptest.NewRequest("GET", "http://localtest/?cep=00000000", nil)
	recorder := httptest.NewRecorder()

	// Act
	handler.GetWeather(recorder, request)
	var response handlers.WeatherResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)

	// Assert
	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, expectedResponse, response)
}

func TestShouldReturnNotFoundIfCEPDoenstHaveInformation(t *testing.T) {
	// Arrange
	cepData := internal.CepData{}

	cepService := mocks.CepServiceMock{}
	cepService.On("GetCepData", mock.AnythingOfType("string")).Return(cepData, nil)
	request := httptest.NewRequest("GET", "http://localtest/?cep=00000000", nil)
	recorder := httptest.NewRecorder()
	handler := handlers.NewWeatherHandler(&cepService, &mocks.WeatherServiceMock{})

	// Act
	handler.GetWeather(recorder, request)

	// Assert
	assert.Equal(t, 404, recorder.Code)
	assert.Equal(t, "can not find zipcode", recorder.Body.String())
}

func TestShouldReturnUnprocessableIfCepIsNotValid(t *testing.T) {
	// Arrange
	request := httptest.NewRequest("GET", "http://localtest/?cep=0", nil)
	recorder := httptest.NewRecorder()
	handler := handlers.NewWeatherHandler(&mocks.CepServiceMock{}, &mocks.WeatherServiceMock{})

	// Act
	handler.GetWeather(recorder, request)

	// Assert
	assert.Equal(t, 422, recorder.Code)
	assert.Equal(t, "invalid zipcode", recorder.Body.String())
}
