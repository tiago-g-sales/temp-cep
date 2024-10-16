package usecase

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
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


func FindCep(cep string) (*ViaCEP, error){
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	var c ViaCEP
	err= json.Unmarshal(body, &c)
	if err != nil{
		return nil, err
	}

	return &c, nil
}