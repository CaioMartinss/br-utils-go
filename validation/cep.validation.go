// Package validation contém funções para validação de dados relacionados a CEP, CPF, CNPJ, etc.
package validation

import (
	"regexp" // Pacote regexp para manipulação de expressões regulares
)

// ValidaCEP verifica se um CEP está no formato válido
func ValidaCEP(cep string) bool {
	// Expressão regular para validar o formato de CEP (opcionalmente com hífen)
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return cepRegex.MatchString(cep)
}
