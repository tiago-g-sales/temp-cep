package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tiago-g-sales/temp-cep/internal/model"
)

type FindCepHTTPClient interface {
    FindCep(cep string) (*model.ViaCEP, error)
}

type HTTPClient struct {
    client *http.Client
}

func NewHTTPClient (client http.Client) (*HTTPClient){
	return &HTTPClient{client: &client}
}


func (h *HTTPClient) FindCep( cep string) (*model.ViaCEP, error){
	
	resp, err := h.client.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	var c model.ViaCEP
	err= json.Unmarshal(body, &c)
	if err != nil{
		return nil, err
	}

	return &c, nil
}