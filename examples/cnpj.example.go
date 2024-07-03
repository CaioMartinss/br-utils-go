package examples

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/validation"

	"github.com/CaioMartinss/br-utils-go/external"
)

func Cnpj_function() {
	cnpj := "00000000000191"
	if !validation.ValidaCNPJ(cnpj) {
		fmt.Println("CNPJ inválido")
		return
	}

	empresa, err := external.ConsultaCNPJ(cnpj)
	if err != nil {
		fmt.Println("Erro ao consultar CNPJ:", err)
		return
	}

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