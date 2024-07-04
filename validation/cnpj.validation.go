// Package validation contém funções para validação e formatação de CNPJ.
package validation

import (
	//"regexp"
	"fmt"     // Pacote fmt para formatação de saída
	"strconv" // Pacote strconv para conversão de tipos
	"strings" // Pacote strings para manipulação de strings
)

// limpa remove caracteres especiais do CNPJ
func Limpa(cnpj string) string {
	return strings.NewReplacer(".", "", "-", "", "/", "").Replace(cnpj)
}

// todosDigitosIguais verifica se todos os dígitos do CNPJ são iguais
func TodosDigitosIguais(cnpj string) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			return false
		}
	}
	return true
}

// calculaDigito calcula o dígito verificador do CNPJ
func CalculaDigito(cnpj string, peso int) int {
	soma := 0
	for i := 0; i < len(cnpj); i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * peso
		peso--
		if peso < 2 {
			peso = 9
		}
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

// ValidaCNPJ verifica se um CNPJ é válido
func ValidaCNPJ(cnpj string) bool {
	cnpj = Limpa(cnpj)
	if len(cnpj) != 14 {
		return false
	}
	if TodosDigitosIguais(cnpj) {
		return false
	}
	d1 := CalculaDigito(cnpj[:12], 5)
	d2 := CalculaDigito(cnpj[:12]+strconv.Itoa(d1), 6)
	return cnpj == cnpj[:12]+strconv.Itoa(d1)+strconv.Itoa(d2)
}

// FormataCNPJ formata um CNPJ com pontos, barras e hífen
func FormataCNPJ(cnpj string) string {
	cnpj = Limpa(cnpj)
	if len(cnpj) != 14 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:])
}
