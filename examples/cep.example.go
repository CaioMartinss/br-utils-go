package examples

import (
	"fmt"

	"github.com/CaioMartinss/br-utils-go/validation"
	
	"github.com/CaioMartinss/br-utils-go/external"
)

func Cep_function() {
	cep := "01001-000" // CEP de exemplo
	if !validation.ValidaCEP(cep) {
		fmt.Println("Formato de CEP inválido.")
		return
	}

	dadosCEP, err := external.ConsultaCEP(cep)
	if err != nil {
		fmt.Println("Erro ao consultar o CEP:", err)
		return
	}

	// Exibe os dados do CEP consultado
	fmt.Println("CEP:", dadosCEP.Cep)
	fmt.Println("Logradouro:", dadosCEP.Logradouro)
	fmt.Println("Complemento:", dadosCEP.Complemento)
	fmt.Println("Bairro:", dadosCEP.Bairro)
	fmt.Println("Localidade:", dadosCEP.Localidade)
	fmt.Println("UF:", dadosCEP.Uf)
	fmt.Println("IBGE:", dadosCEP.Ibge)
	fmt.Println("GIA:", dadosCEP.Gia)
	fmt.Println("IBGE:", dadosCEP.Ddd)
	fmt.Println("GIA:", dadosCEP.Siafi)

}