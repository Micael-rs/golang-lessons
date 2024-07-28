package capitulo2

import "fmt"

func Exercicio1() {

	//Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

	valor := 1
	fmt.Printf("%d, %b, %#x", valor, valor, valor)

}

func Exercicio2(){

	//- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
	// Demonstre estes valores.

	a := 10==10

	fmt.Println(a)

}

func Exercicio3(){

	//- Crie constantes tipadas e não-tipadas.
	//- Demonstre seus valores.

	const a = 10
	const b float64 = 10 

	fmt.Println(a, b)
}

func Exercicio4(){

	//- Crie um programa que:
    //- Atribua um valor int a uma variável
    //- Demonstre este valor em decimal, binário e hexadecimal
    //- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    //- Demonstre esta outra variável em decimal, binário e hexadecimal

	a := 200
	fmt.Printf("%d, %b, %#x\n", a, a, a)

	b := a<<1
	fmt.Printf("%d, %b, %#x\n", b, b, b)

}

func Exercicio5(){

	//- Crie uma variável de tipo string utilizando uma raw string literal.
	//- Demonstre-a.

	a := `isto
	         é 
			        uma coisa
					           doida`

	fmt.Println(a)

}

func Exercicio6(){

	const(
		a = 2025 + iota
		b 
		c
		d
	)

	fmt.Println(a, b, c, d)

}