package examples

import (
	"fmt"

	"github.com/CaioMartinss/br-utils-go/validation"
)

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
