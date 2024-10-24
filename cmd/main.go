package main

import (

	"net/http"
	"github.com/tiago-g-sales/temp-cep/configs"
	"github.com/tiago-g-sales/temp-cep/internal/handler"
)



func main(){

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	
	http.HandleFunc("/", handler.FindTempByCepHandler)
	http.ListenAndServe(configs.WebServerPort, nil)
}








