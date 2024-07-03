package validation

import (
	"fmt"
	"strconv"
)

// ValidaCPF verifica se um CPF é válido
func ValidaCPF(cpf string) bool {
	cpf = limpa(cpf)
	if len(cpf) != 11 {
		return false
	}

	if todosDigitosIguais(cpf) {
		return false
	}

	d1 := calculaDigito(cpf[:9], 10)
	d2 := calculaDigito(cpf[:9]+strconv.Itoa(d1), 11)

	return cpf == cpf[:9]+strconv.Itoa(d1)+strconv.Itoa(d2)
}

// FormataCPF formata um CPF com pontos e hífen
func FormataCPF(cpf string) string {
	cpf = limpa(cpf)
	if len(cpf) != 11 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s-%s", cpf[:3], cpf[3:6], cpf[6:9], cpf[9:])
}