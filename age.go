// Informe se é menor ou maior de idade

package main

import (
	"fmt"
	"time"
)

func idade(a int) int {
	i := time.Now().Year() - a
	return i
}

func main() {

	// Pegar o valor da idade
	fmt.Println("Informe seu ano de nascimento: ")
	var i int
	fmt.Scanln(&i)

	// retorna se é ou não maior de idade

	idade := idade(i)
	fmt.Printf("Sua idade é %d anos", idade)

	if idade < 18 {
		fmt.Println("\nVocê é menor de idade!")
	} else {
		fmt.Println("\nVocê é maior de idade!")
	}
}