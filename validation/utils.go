// Package validation contém funções auxiliares para validação e formatação de dados numéricos.
package validation

import (
	"strconv" // Pacote strconv para conversão de tipos
	"strings" // Pacote strings para manipulação de strings
)

// limpa remove caracteres não numéricos de uma string
func limpa(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// todosDigitosIguais verifica se todos os dígitos de uma string são iguais
func todosDigitosIguais(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

// calculaDigito calcula o dígito de verificação de uma string
func calculaDigito(s string, pesoInicial int) int {
	soma := 0
	peso := pesoInicial
	for _, r := range s {
		num, _ := strconv.Atoi(string(r))
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
