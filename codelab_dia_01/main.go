package main

import "fmt"

func main() {

// Exercícios dia 1

		fmt.Println("<< Exercício #1 >>")

	var num1 = 21
	var num2 = 28
	var num3 = 8

	fmt.Println("Num1:", num1)
	fmt.Println("Num2:", num2)
	fmt.Println("Num3:", num3)

		fmt.Println("\n<< Exercício #2 >>")

	a := 230
	b := 27

	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("SumAB=", a+b)
	fmt.Println("SubAB=", a-b)

		fmt.Println("\n<< Exercício #3 >>")

	var div = float64(a) / float64(b)

	fmt.Println("MultAB=", a*b)
	fmt.Println("DivAB=", div)

		fmt.Println("\n<< Exercício #4 >>")
	
	fmt.Printf("A e B são variáveis: %T e %T", a, b)
	fmt.Printf("\nSumAB e SubAB são variáveis: %T e %T", a+b, a-b)
	fmt.Printf("\nMultAB e DivAB são variáveis: %T e %T\n", a*b, div)

		fmt.Println("\n<< Exercício #5 >>")
	//Preços dos itens do mercado

	var preçoDoPão float64 = 4.60
	var preçoDaAveia float64 = 5
	var preçoDoAzeite float64 = 19.95
	var preçoDoSuco float64 = 7
	var preçoDaÁgua float64 = 2.15
	var preçoDoKGMaçã float64 = 8.90
	var preçoDoKGBanana float64 = 4.5
/*
3 pães
1 azeite
2 garrafas de suco de laranja
5 garrafas de água
1.5 kg de maçã
*/

	TotalPão := preçoDoPão * 3
	TotalAzeite := preçoDoAzeite * 1
	TotalSuco := preçoDoSuco * 2
	TotalAgua := preçoDaÁgua * 5
	TotalMaça := preçoDoKGMaçã * 1.5
	TotalBanana := preçoDoKGBanana * 0
	TotalAveia := preçoDaAveia * 0
	ValorDaCompra := preçoDoPão * 3 +
	preçoDoAzeite * 1 +
	preçoDoSuco * 2 +
	preçoDaÁgua * 5 +
	preçoDoKGMaçã * 1.5 +
	preçoDaAveia * 0 +
	preçoDoKGBanana * 0
	
	fmt.Printf("::Total Pão:R$ %.2f", TotalPão)
	fmt.Printf("\n::Total Suco:R$ %.2f", TotalSuco)
	fmt.Printf("\n::Total Maçã:R$ %.2f", TotalMaça)
	fmt.Println("\n::Total Azeite:R$", TotalAzeite ,"\n::Total Água:R$", TotalAgua ,"\n::Total Banana:R$", TotalBanana ,"\n::Total Aveia:R$", TotalAveia)
	fmt.Println("Valor total da Compra:R$", ValorDaCompra)

		fmt.Println("\n<< Exercício 6 >>")
// var := 27
// fmt.Println("A:", var)
	fmt.Println("A palavra 'var' é uma keyword para declarar uma variável e não pode ser usada como identificador")
		
		fmt.Println("\n<< Exercício 7 >>")

	x := 15
	var y int = 31
	var z = 47

	fmt.Printf("\t X\t Y\t Z")
	fmt.Printf("\nBase 2: %b \t%b \t%b", x, y, z)
	fmt.Printf("\nBase 10: %d \t%d \t%b", x, y, z)
	fmt.Printf("\nBase 16: %x \t%x \t%x", x, y, z)
	
	
		fmt.Println("\n\n<< Exercício #8 >>")
	
		var X int
		
		fmt.Println("Valor de X:", X)
		fmt.Println("O 0 'zero' é o valor default para variável do tipo int. Portanto, caso nenhum valor seja atribuído, o valor '0' será considerado")

		
}