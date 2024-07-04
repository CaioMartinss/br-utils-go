// Package examples contém exemplos de uso das funções de validação e consulta de CEP.
package examples

import (
	"fmt" // Pacote fmt para formatação de saída
	"github.com/CaioMartinss/br-utils-go/validation" // Pacote validation para validação de CEP
	"github.com/CaioMartinss/br-utils-go/external"   // Pacote external para consulta de CEP
)

// Cep_function demonstra a validação e consulta de um CEP no Brasil.
// A função valida o formato do CEP usando a função ValidaCEP do pacote validation.
// Em seguida, consulta os dados do CEP usando a função ConsultaCEP do pacote external.
// Os dados do CEP são exibidos na saída padrão. Em caso de formato inválido do CEP
// ou erro na consulta, mensagens apropriadas são exibidas.
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
	fmt.Println("DDD:", dadosCEP.Ddd)
	fmt.Println("SIAFI:", dadosCEP.Siafi)
}
