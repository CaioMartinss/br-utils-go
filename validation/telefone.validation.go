package validation

import "fmt"

// ValidaNumero valida se um número está no formato correto de um número Telefonico
func ValidaNumero(numero string) bool {
	numero = limpa(numero)
	if len(numero) != 10 {
		return false
	}

	if todosDigitosIguais(numero) {
		return false
	}

	return true
}

// FormataNumero formata um Numero , com o DDD entre parênteses
func FormataNumero(numero string) string {
	numero = limpa(numero)
	if len(numero) != 10 {
		return ""
	}

	return fmt.Sprintf("(%s)%s", numero[:2], numero[2:])
}

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

func DDDporEstado(ddd string) string {
	var estado string
	ddd = limpa(ddd)
	switch ddd {
	case "11":
		estado = "São Paulo"
	case "21":
		estado = "Rio de Janeiro"
	case "31":
		estado = "Minas Gerais"
	case "41":
		estado = "Paraná"
	case "51":
		estado = "Rio Grande do Sul"
	case "61":
		estado = "Distrito Federal"
	case "71":
		estado = "Bahia"
	case "81", "87":
		estado = "Pernambuco"
	case "91", "93", "94":
		estado = "Pará"
	case "62", "64":
		estado = "Goiás"
	case "65", "66":
		estado = "Mato Grosso"
	case "67":
		estado = "Mato Grosso do Sul"
	case "82":
		estado = "Alagoas"
	case "83":
		estado = "Paraíba"
	case "84":
		estado = "Rio Grande do Norte"
	case "85", "88":
		estado = "Ceará"
	case "86", "89":
		estado = "Piauí"
	case "92", "97":
		estado = "Amazonas"
	case "95":
		estado = "Roraima"
	case "96":
		estado = "Amapá"
	case "98", "99":
		estado = "Maranhão"
	default:
		estado = "DDD não reconhecido"
	}

	return estado
}
