package mocks

import (
	"github.com/Survialander/weather/internal"
	"github.com/stretchr/testify/mock"
)

type CepServiceMock struct {
	mock.Mock
}

func (m *CepServiceMock) GetCepData(cep string) (internal.CepData, error) {
	args := m.Called(cep)
	return args.Get(0).(internal.CepData), args.Error(1)
}
