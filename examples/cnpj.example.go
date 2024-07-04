// Package examples contém exemplos de uso das funções de validação e consulta de CNPJ.
package examples

import (
	"fmt" // Pacote fmt para formatação de saída
	"github.com/CaioMartinss/br-utils-go/validation" // Pacote validation para validação e formatação de CNPJ
	"github.com/CaioMartinss/br-utils-go/external"   // Pacote external para consulta de CNPJ
)

// Cnpj_function demonstra a validação e consulta de um CNPJ no Brasil.
// A função valida o formato do CNPJ usando a função ValidaCNPJ do pacote validation.
// Em seguida, consulta os dados do CNPJ usando a função ConsultaCNPJ do pacote external.
// Os dados do CNPJ são formatados e exibidos na saída padrão. Em caso de formato inválido do CNPJ
// ou erro na consulta, mensagens apropriadas são exibidas.
func Cnpj_function() {
	cnpj := "00000000000191" // CNPJ de exemplo
	if !validation.ValidaCNPJ(cnpj) {
		fmt.Println("CNPJ inválido")
		return
	}

	empresa, err := external.ConsultaCNPJ(cnpj)
	if err != nil {
		fmt.Println("Erro ao consultar CNPJ:", err)
		return
	}

	// Exibe os dados do CNPJ consultado
	fmt.Printf("CNPJ: %s\n", validation.FormataCNPJ(cnpj))
	fmt.Printf("Nome: %s\n", empresa.Nome)
	fmt.Printf("Nome Fantasia: %s\n", empresa.Fantasia)
	if len(empresa.Atividades) > 0 {
		fmt.Printf("Atividade Principal: %s\n", empresa.Atividades[0].Texto)
	}
	fmt.Printf("Município: %s\n", empresa.Municipio)
	fmt.Printf("UF: %s\n", empresa.UF)
	fmt.Printf("Status: %s\n", empresa.Status)
}
