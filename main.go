package main

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/cmd"
)

func main() {
	numero := "(11)912345678"
	numero2 := "11912345678"
	numero3 := "1191234567"
	fmt.Printf("numero %s é válido? %v\n", numero, brutils.ValidaNumeroMovel(numero))
	fmt.Printf("numero2 %s é válido? %v\n", numero2, brutils.ValidaNumeroMovel(numero2))
	fmt.Printf("numero3 %s é válido? %v\n", numero3, brutils.ValidaNumeroMovel(numero3))
}
