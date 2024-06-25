package brutils

import (
	"regexp"
)

// ValidaNumeroMovel valida se um número está no formato correto de um número móvel
func ValidaNumeroMovel(numero string) bool {
	// Expressão regular para validar o formato de um número móvel
	re := regexp.MustCompile(`^\(?\d{2}\)?9\d{8}$`)
	return re.MatchString(numero)
}

// FormataNumero verifica se um número de telefone está no formato correto.
func FormataNumero(numero string) bool {
	numero = limpa(numero)
	// Expressão regular para validar o formato do número de telefone
	padrao := `^\(\d{2}\)\s?\d{4,5}-\d{4}$`
	regex := regexp.MustCompile(padrao)
	return regex.MatchString(numero)
}
