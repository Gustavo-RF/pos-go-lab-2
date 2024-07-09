package entities

type ZipCodeResponse struct {
	Localidade string `json:"localidade"`
}

func NewZipCodeResponse(localidade string) ZipCodeResponse {
	return ZipCodeResponse{
		Localidade: localidade,
	}
}
