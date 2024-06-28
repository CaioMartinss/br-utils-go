package main

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/cmd"
)

func main() {
	cpf := "123.456.789-09"
	fmt.Printf("CPF %s é válido? %v\n", cpf, brutils.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", brutils.FormataCPF(cpf))

	// cnpj := "12.345.678/0001-95"
	// fmt.Printf("CNPJ %s é válido? %v\n", cnpj, brutils.ValidaCNPJ(cnpj))
	// fmt.Printf("CNPJ formatado: %s\n", brutils.FormataCNPJ(cnpj))

	cep := "01001-000" // CEP de exemplo
	if !brutils.ValidaCEP(cep) {
		fmt.Println("Formato de CEP inválido.")
		return
	}

	dadosCEP, err := brutils.ConsultaCEP(cep)
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


	cnpj := "00000000000191"
	if !brutils.ValidaCNPJ(cnpj) {
		fmt.Println("CNPJ inválido")
		return
	}

	empresa, err := brutils.ConsultaCNPJ(cnpj)
	if err != nil {
		fmt.Println("Erro ao consultar CNPJ:", err)
		return
	}

	fmt.Printf("CNPJ: %s\n", brutils.FormataCNPJ(cnpj))
	fmt.Printf("Nome: %s\n", empresa.Nome)
	fmt.Printf("Nome Fantasia: %s\n", empresa.Fantasia)
	if len(empresa.Atividades) > 0 {
		fmt.Printf("Atividade Principal: %s\n", empresa.Atividades[0].Texto)
	}
	fmt.Printf("Município: %s\n", empresa.Municipio)
	fmt.Printf("UF: %s\n", empresa.UF)
	fmt.Printf("Status: %s\n", empresa.Status)
}