package main

import "fmt"


func main(){
	fmt.Println("Digite seu nome: ")
	var nome string
	fmt.Scanln(&nome)
	fmt.Printf("Olá \033[1;32m%s\033[m, seja bem-vindo!\nEu sou a linguagem \033[1;34mGO\033[m!!!", nome)
}