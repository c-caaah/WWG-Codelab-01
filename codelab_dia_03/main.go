package main

import "fmt"

func mais(a int, b int, c int) int {

	return a + b + c

}

func MultiCinc(num int) {
	if num%5 == 0 {
		fmt.Println("É multiplo de 5")
	} else {
		fmt.Println("Não é multiplo de 5")
	}
}

func main() {

	// Exercícios dia 3

	fmt.Println("<< Exercício #1 >>")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("\n<< Exercício #2 >>")
	for i2 := 1; i2 <= 100; i2++ {
		if i2%9 != 0 {
			continue

		}
		fmt.Println(i2)

	}

	fmt.Println("\n<< Exercício #3 >>")
	MultiCinc(259)

	fmt.Println("\n<< Exercício #4 >>")

	res := mais(2, 5, 7)

	fmt.Println("Soma", res)

}
