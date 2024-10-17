package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/tiago-g-sales/temp-cep/configs"
	"github.com/tiago-g-sales/temp-cep/pkg"
	"github.com/valyala/fastjson"
)

type Temperatura struct {
	Temp_C float64 `json:"temp_C"`
	Temp_F float64 `json:"temp_F"`
	Temp_K float64 `json:"temp_K"`
}

func FindTemp(loc string) (*Temperatura, error) {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	localidade := pkg.Replace(loc)

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

	temp := Temperatura{}
	json.Unmarshal([]byte(v.GetObject("current").String()), &temp)

	temp.Temp_K = pkg.ConvertTemp(temp.Temp_C)

	return &temp, nil
}
