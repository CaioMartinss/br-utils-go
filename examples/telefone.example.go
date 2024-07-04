// Package examples contém exemplos de uso das funções de validação e formatação de números de telefone.
package examples

import (
	"fmt" // Pacote fmt para formatação de saída
	"github.com/CaioMartinss/br-utils-go/validation" // Pacote validation para validação e formatação de números de telefone
)

// Telefone_function demonstra a validação e formatação de números de telefone no Brasil.
// A função mostra como validar e formatar números de telefone móveis e fixos usando as funções
// ValidaNumeroMovel, FormataNumeroMovel, ValidaNumero e FormataNumero do pacote validation.
// Além disso, a função demonstra como obter o estado brasileiro a partir do DDD usando a função DDDporEstado.
func Telefone_function() {
	numeroMovel := "11111111213"
	fmt.Println(validation.ValidaNumeroMovel(numeroMovel))
	fmt.Println(validation.FormataNumeroMovel(numeroMovel))

	numero := "(78)34864590"
	fmt.Println(validation.ValidaNumero(numero))
	fmt.Println(validation.FormataNumero(numero))

	ddd := "31"
	fmt.Println(validation.DDDporEstado(ddd))
	ddd2 := "94"
	fmt.Println(validation.DDDporEstado(ddd2))
}
