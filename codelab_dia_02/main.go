package main

import "fmt"

func main() {
	// Exercícios dia 02

	fmt.Println("<< Exercício #1 >>")
	x := 12
	y := 80

	fmt.Println("O Maior valor é:")
	if x > y {
		fmt.Printf("y = %v", x)
	} else if x < y {
		fmt.Printf("x = %v", y)
	} else {
		fmt.Printf("os valores são iguais")
	}

	fmt.Println("\n\n<< Exercício #2 >>")
	Ano := 1991

	if Ano > (2019 - 16) {
		fmt.Println("Não está apto")
	} else {
		fmt.Println("Está apto")
	}

	fmt.Println("\n<< Exercício #3 >>")

	senha := "abc@123"
	entrada := "ABC@12"

	if senha == entrada {
		fmt.Println("Acesso Permitido")
	} else {
		fmt.Println("Acesso Negado")
	}

	fmt.Println("\n<< Desafio >>")

	variavel := "Hello Aliens"
	tipoVariavel := fmt.Sprintf("%T", variavel)

	fmt.Println("Esta variável é:", tipoVariavel)
}
