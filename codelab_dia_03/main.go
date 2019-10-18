package main

import "fmt"

func main() {

	// Exercícios dia 3

	fmt.Println("<< Exercício #1 >>")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("\n<< Exercício #2 >>")
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\n<< Exercício #3 >>")
	var num = 259
	if num%5 == 0 {
		fmt.Println("É multiplo de 5")
	} else {
		fmt.Println("Não é multiplo de 5")
	}

}
