package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiago-g-sales/temp-cep/internal/usecase"
	"github.com/tiago-g-sales/temp-cep/internal/usecase/mocks"
	"github.com/tiago-g-sales/temp-cep/internal/usecase/model"
)


func TestFindCep(t *testing.T ){

	endereco :=  model.ViaCEP{
		Cep: "06865-010",
		Logradouro: "Rua Tancredo de Almeida Neves",
		Complemento: "",
		Bairro: "Jardim do Édem",
		Localidade: "Itapecerica da Serra",
		Uf: "SP",
		Ibge: "3522208",
		Gia: "3700",
		Ddd: "11",
		Siafi: "6545",
	}



	c := mocks.NewMockClientCep()
	c.On("FindCep", endereco.Cep,).Return(&model.ViaCEP{
		Cep: endereco.Cep,
		Logradouro: "Rua Tancredo de Almeida Neves",
		Complemento: "",
		Bairro: "Jardim do Édem",
		Localidade: "Itapecerica da Serra",
		Uf: "SP",
		Ibge: "3522208",
		Gia: "3700",
		Ddd: "11",
		Siafi: "6545",
	}, nil)

	temp , err := usecase.FindCepHTTPClient.FindCep(c, endereco.Cep)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, temp.Cep, endereco.Cep )
	assert.Equal(t, temp.Logradouro, endereco.Logradouro )
	assert.Equal(t, temp.Bairro, endereco.Bairro)
	assert.Equal(t, temp.Complemento, endereco.Complemento)
	assert.Equal(t, temp.Ddd, endereco.Ddd)
	assert.Equal(t, temp.Gia, endereco.Gia)
	assert.Equal(t, temp.Ibge, endereco.Ibge)
	assert.Equal(t, temp.Localidade, endereco.Localidade)
	assert.Equal(t, temp.Uf, endereco.Uf)
	assert.Equal(t, temp.Siafi, endereco.Siafi)

}