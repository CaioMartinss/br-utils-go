package brutils

import (
	"fmt"
)

// ValidaNumeroMovel valida se um número está no formato correto de um número móvel
func ValidaNumeroMovel(numero string) bool {
	numero = limpa(numero)
	if len(numero) != 11 {
		return false
	}

	if todosDigitosIguais(numero) {
		return false
	}

	return true
}

// FormataNumeroMovel formata um Numero Movel, com o DDD entre parênteses
func FormataNumeroMovel(numero string) string {
	numero = limpa(numero)
	if len(numero) != 11 {
		return ""
	}

	return fmt.Sprintf("(%s)%s", numero[:2], numero[2:])
}

