package main

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/pacotes/brutils"
)

func main() {
	cpf := "123.456.789-09"
	fmt.Printf("CPF %s é válido? %v\n", cpf, brutils.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", brutils.FormataCPF(cpf))

	cnpj := "12.345.678/0001-95"
	fmt.Printf("CNPJ %s é válido? %v\n", cnpj, brutils.ValidaCNPJ(cnpj))
	fmt.Printf("CNPJ formatado: %s\n", brutils.FormataCNPJ(cnpj))
}
