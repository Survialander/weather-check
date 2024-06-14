package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Survialander/weather/internal/utils"
)

type CepService struct{}

type CepServiceInterface interface {
	GetCepData(cep string) (CepData, error)
}

type CepData struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	Uf          string
}

func NewCepService() *CepService {
	return &CepService{}
}

func (s *CepService) GetCepData(cep string) (CepData, error) {
	client := utils.GetHttpClient()

	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cep)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := client.Do(req)

	if err != nil {
		return CepData{}, err
	}

	var data CepData
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return CepData{}, err
	}

	return data, nil
}
