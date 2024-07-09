package entities

import "encoding/json"

type ZipCodeApiResponse struct {
	Erro        string `json:"erro"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewZipCodeApiResponse(zipCode []byte) (*ZipCodeApiResponse, error) {
	var data ZipCodeApiResponse
	err := json.Unmarshal(zipCode, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
