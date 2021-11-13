// Funções em golang que realiza as 4 operações

package main

import "fmt"

func Soma(x int, y int) int {
	return x + y
} 

func Subtração(x int, y int) int {
	return x - y
}

func Multiplicação(x int, y int) int {
	return x * y
}

func Divisão(x int, y int) int {
	return x / y
}

func main() {

	a := 100
	b := 20

	c := Soma(a, b)
	d := Subtração(a, b)
	e := Multiplicação(a, b)
	f := Divisão(a, b)

	fmt.Println("\033[1;32mSoma\033[m:\nO resultado da operação de soma é:", c)
	fmt.Println("\033[1;32mSubtração\033[m:\nO resultado da operação de subtração é:", d)
	fmt.Println("\033[1;32mMultiplicação\033[m\nO resultado da operação é:", e)
	fmt.Println("\033[1;32mDivisão\033[m\nO resultado da operação de divisão é:", f)
}