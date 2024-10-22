package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tiago-g-sales/temp-cep/configs"
	"github.com/tiago-g-sales/temp-cep/internal/usecase/model"
	pkg_test "github.com/tiago-g-sales/temp-cep/pkg"

	"github.com/valyala/fastjson"
)

type FindTempHTTPClient interface {
    FindTemp(loc string) (*model.Temperatura, error)
}

type HTTPClientTemp struct {
    client *http.Client
}

func NewHTTPClientTemp (client http.Client) (*HTTPClientTemp){
	return &HTTPClientTemp{client: &client}
}


func (h *HTTPClientTemp) FindTemp(loc string) (*model.Temperatura, error) {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	localidade := pkg_test.ReplaceAndRemoveAccents(loc)

	resp, err := http.Get(fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&lang=json&key=%s",localidade , configs.API_KEY))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p fastjson.Parser

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	v, err := p.ParseBytes(body)
	if err != nil {
		panic(err)
	}

	temp := model.Temperatura{}
	json.Unmarshal([]byte(v.GetObject("current").String()), &temp)

	temp.Temp_K, _ = pkg_test.ConvertTemp(temp.Temp_C)

	return &temp, nil
}
