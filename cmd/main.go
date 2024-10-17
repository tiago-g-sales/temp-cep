package main

import (
	"encoding/json"
	"net/http"

	"github.com/tiago-g-sales/temp-cep/configs"
	"github.com/tiago-g-sales/temp-cep/internal/usecase"
)

const(
	INVALID_ZIP_CODE = "invalid zipcode"
	CAN_NOT_FIND_ZIPCODE = "can not find zipcode"
	QUERY_PARAMETER = "cep"
	LEN_ZIP_CODE = 8
)



func main(){

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(configs.WebServerPort, nil)
}


func BuscaCepHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get(QUERY_PARAMETER)
	if cepParam == ""{
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(cepParam) > LEN_ZIP_CODE || len(cepParam) < LEN_ZIP_CODE  {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(INVALID_ZIP_CODE)
		return		
	}

	cep, err :=usecase.FindCep(cepParam)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(CAN_NOT_FIND_ZIPCODE)
		return
	}

	temp , err := usecase.FindTemp(cep.Localidade)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return		
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)

}






