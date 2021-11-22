package main

import "fmt"

func main() {

	var(
		peso float64 = 80.0
		idade int = 19
		altura float64 = 1.80
	)

	fmt.Println("Olá")
	fmt.Println("Sua idade é igual a:", idade)
	fmt.Printf("Você apresentou um peso de %f kg e uma altura de %f metros", peso, altura)
}