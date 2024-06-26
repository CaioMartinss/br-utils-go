package main

import (
	"fmt"

	brutils "github.com/CaioMartinss/br-utils-go/cmd"
)

func main() {
	numero := "11111111213"
	fmt.Println(brutils.ValidaNumeroMovel(numero))
	fmt.Println(brutils.FormataNumeroMovel(numero))

}
