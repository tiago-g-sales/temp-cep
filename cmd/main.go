package main

import (
	"encoding/json"
	"net/http"
	"github.com/tiago-g-sales/temp-cep/internal/usecase"

)




func main(){

	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}


func BuscaCepHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == ""{
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(cepParam) > 8 || len(cepParam) < 8  {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid zipcode")
		return		
	}

	cep, err :=usecase.FindCep(cepParam)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("can not find zipcode")
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







