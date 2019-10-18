package main

import "fmt"

func main() {
	// Exercícios dia 4

	fmt.Println("<< Exercício #1 >>")
	fmt.Println("\nArray Int")
	a := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	fmt.Println("Resultado:", a)

	/* for range
	   por parágrafo
	*/
	fmt.Println("\nFor Range")
	for v := range a {
		fmt.Println(v)

	}

	fmt.Println("\n<< Exercício #2 >>")

	numeros := []int{40, 75, 157}
	soma(numeros)
}

func soma(operandos []int) {
	soma := 0
	for _, valorOperando := range operandos {
		soma = soma + valorOperando
	}

	fmt.Println(soma)
}
