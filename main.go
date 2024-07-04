// Package main é o ponto de entrada do programa.
package main

import (
	"fmt"

	"github.com/CaioMartinss/br-utils-go/examples"
)

func main() {
	// Imprime um exemplo de validação de CPF
	fmt.Println("Exemplo de validação de CPF:")
	examples.Cpf_function()

	// Imprime um exemplo de validação de CNPJ
	fmt.Println("\nExemplo de validação de CNPJ:")
	examples.Cnpj_function()

	// Imprime um exemplo de validação de CEP
	fmt.Println("\nExemplo de validação de CEP:")
	examples.Cep_function()

	// Imprime um exemplo de validação de Telefone
	fmt.Println("\nExemplo de validação de Telefone:")
	examples.Telefone_function()
}
