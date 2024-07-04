package main

import (
	"fmt"

	"github.com/CaioMartinss/br-utils-go/examples"
)

func main() {
	fmt.Println("Exemplo de validação de CPF:")
	examples.Cpf_function()

	fmt.Println("\nExemplo de validação de CNPJ:")
	examples.Cnpj_function()

	fmt.Println("\nExemplo de validação de CEP:")
	examples.Cep_function()

	fmt.Println("\nExemplo de validação de Telefone:")
	examples.Telefone_function()
}
