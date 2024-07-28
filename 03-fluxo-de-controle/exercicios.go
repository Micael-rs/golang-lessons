package capitulo3

import "fmt"

func Desafio() {

	//- Desafio surpresa!
	//- Format printing:
	//- Decimal       %d
	//- Hexadecimal   %#x
	//- Unicode       %#U
	//- Tab           \t
	//- Linha nova    \n
	//- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.

	for contador := 33; contador <= 122; contador++ {
		fmt.Printf("%c\n", contador)
	}

}

func Exercicio1() {

	// Põe na tela: todos os números de 1 a 10000.

	for contador := 1; contador <= 10000; contador++ {
		fmt.Print(contador, " ")
	}

}

func Exercicio2() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}

func Exercicio3() {

	a := 2006
	b := 2024

	for a <= b {
		fmt.Println(a)
		a++
	}

}

func Exercicio4() {

	//- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}

}
