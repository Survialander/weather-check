package mocks

import (
	"github.com/Survialander/weather/internal"
	"github.com/stretchr/testify/mock"
)

type WeatherServiceMock struct {
	mock.Mock
}

func (m *WeatherServiceMock) GetWeatherData(city string) (internal.WeatherData, error) {
	args := m.Called(city)
	return args.Get(0).(internal.WeatherData), args.Error(1)
}
