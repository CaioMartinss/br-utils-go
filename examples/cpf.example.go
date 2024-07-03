package examples

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/validation"
)

func Cpf_function() {
	cpf := "123.456.789-09"
	fmt.Printf("CPF %s é válido? %v\n", cpf, validation.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", validation.FormataCPF(cpf))

	cpf = "12345678909"
	fmt.Printf("CPF %s é válido? %v\n", cpf, validation.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", validation.FormataCPF(cpf))


}