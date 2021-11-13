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

	// Informar o ano de nascimento
	fmt.Println("Informe seu ano de nascimento:")
	var ano_n int
	fmt.Scanln(&ano_n)

	idade := idade(ano_n)
	fmt.Printf("\nIdade: %d anos", idade)

	// Condições
	if idade < 18 {
		fmt.Println("\nVocê é uma criança.")
	} else {
		if idade >= 18 && idade < 24 {
			fmt.Println("\nVocê é um jovem.")
		} else {
			if idade >= 24 && idade < 45 {
				fmt.Println("\nVocê é um adulto.")
			} else {
				if idade >= 45 && idade < 60 {
					fmt.Println("\nVocê é uma pessoa de meia-idade.")
				} else {
					if idade >= 60 {
						fmt.Println("\nVocê é um idoso.")
					}
				}
			}
		}
	}

	fmt.Println("\nObrigado por participar!")
}