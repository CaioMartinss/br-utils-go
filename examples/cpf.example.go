// Package examples contém exemplos de uso das funções de validação e formatação de CPF.
package examples

import (
	"fmt" // Pacote fmt para formatação de saída
	"github.com/CaioMartinss/br-utils-go/validation" // Pacote validation para validação e formatação de CPF
)

// Cpf_function demonstra a validação e formatação de CPFs no Brasil.
// A função mostra como validar e formatar CPFs usando as funções ValidaCPF e FormataCPF do pacote validation.
func Cpf_function() {
	cpf := "123.456.789-09"
	fmt.Printf("CPF %s é válido? %v\n", cpf, validation.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", validation.FormataCPF(cpf))

	cpf = "12345678909"
	fmt.Printf("CPF %s é válido? %v\n", cpf, validation.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", validation.FormataCPF(cpf))
}
