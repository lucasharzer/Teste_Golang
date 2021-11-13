package main

import (
	"fmt"
	"time"
)

func Idade(n int) int {
	i := time.Now().Year() - n
	return i
}

func main() {

	// Nome
	fmt.Println("Qual é o seu nome? ")
	var nome string
	fmt.Scanln(&nome)

	// Idade
	fmt.Println("Qual o seu ano de nascimento? ")
	var ano_n int
	fmt.Scanln(&ano_n)

	// Mensagem
	fmt.Println("Sua profissão: ")
	var pf string
	fmt.Scanln(&pf)

	fmt.Println("-----------------------")

	fmt.Println("\033[1;34mInformações:\033[m")
	fmt.Printf("Nome: %s", nome)
	fmt.Println("\nIdade:", Idade(ano_n), "anos")
	fmt.Println("Profissão:", pf)
	
	fmt.Println("\033[1;34mObrigado por participar!!!\033[m")
}