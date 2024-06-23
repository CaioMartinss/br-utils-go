package brutils

import (
	"fmt"
	"strconv"
)

// ValidaCNPJ verifica se um CNPJ é válido
func ValidaCNPJ(cnpj string) bool {
	cnpj = limpa(cnpj)
	if len(cnpj) != 14 {
		return false
	}

	if todosDigitosIguais(cnpj) {
		return false
	}

	d1 := calculaDigito(cnpj[:12], 5)
	d2 := calculaDigito(cnpj[:12]+strconv.Itoa(d1), 6)

	return cnpj == cnpj[:12]+strconv.Itoa(d1)+strconv.Itoa(d2)
}

// FormataCNPJ formata um CNPJ com pontos, barras e hífen
func FormataCNPJ(cnpj string) string {
	cnpj = limpa(cnpj)
	if len(cnpj) != 14 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:])
}