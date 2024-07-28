package capitulo6

import (
	"fmt"
	"math"
)

//-----------------------------------------------

func Exercicio1() {

	//- Crie uma função que retorne um int
	//- Crie outra função que retorne um int e uma string
	//- Chame as duas funções
	//- Demonstre seus resultados

	fmt.Println(inteiro())
	fmt.Println(texto())

}

func inteiro() int {
	a := 1
	return a
}

func texto() (string, int) {
	a := "Oi"
	b := 2
	return a, b
}

//----------------------------------------------------

func Exercicio2() {

	//- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
	//- Passe um valor do tipo slice of int como argumento para a função;
	//- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
	//- Passe um valor do tipo slice of int como argumento para a função.

	a := variaticos(10, 20, 30, 40, 50)
	b := []int{10, 20, 30, 40, 50}
	c := slice(b...)

	fmt.Println(a)
	fmt.Println(c)
}

func variaticos(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func slice(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

//-----------------------------------------------------------------

func Exercicio3() {

	//- Crie um tipo struct "pessoa" que contenha os campos:
	//- nome
	//- sobrenome
	//- idade
	//- Crie um método para "pessoa" que demonstre o nome completo e a idade;
	//- Crie um valor de tipo "pessoa";
	//- Utilize o método criado para demonstrar esse valor.

	// struct pessoa já existe no package capitulo6

	a := pessoa{"Micael", 18}
	a.mensagem()

}

func (p pessoa) mensagem() {
	fmt.Println("Nome do usuário:", p.nome)
}

//------------------------------------------------------------------------------

func Exercicio4() {

	//- Crie um tipo "quadrado"
	//- Crie um tipo "círculo"
	//- Crie um método "área" para cada tipo que calcule e retorne a área da figura
	//- Área do círculo: 2 * π * raio
	//- Área do quadrado: lado * lado
	//- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
	//- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
	//- Crie um valor de tipo "quadrado"
	//- Crie um valor de tipo "círculo"
	//- Use a função "info" para demonstrar a área do "quadrado"
	//- Use a função "info" para demonstrar a área do "círculo"

	c := circulu{raio: 5}
	t := quadrado{lado: 10}

	info(c)
	info(t)

}

type quadrado struct {
	lado float64
}

type circulu struct {
	raio float64
}

func (a circulu) area() float64 {
	return math.Pi * a.raio * a.raio
}

func (a quadrado) area() float64 {
	return a.lado * a.lado
}

type figura interface {
	area() float64
}

func info(a figura) {
	fmt.Println(a.area())
}

//----------------------------------------------------------------------------------

func Exercicio5() {

	a := func() int {
		x := 10
		return x
	}

	fmt.Println(a())
}

//----------------------------------------------------------------------------------

func Exercicio6() {

	//- Crie uma função que retorna uma função.
	//- Atribua a função retornada a uma variável.
	//- Chame a função retornada.

	a := a()()

	fmt.Println(a)

}

func a() func() int {
	return func() int {
		a := 10
		return a
	}
}

//----------------------------------------------------

func Exercicio7() {

	//- Callback: passe uma função como argumento a outra função.

	essavaireceberaoutrafuncao(issovaiserumargumento)
}

func issovaiserumargumento() {
	fmt.Println("Olha eu aqui!")
}

func essavaireceberaoutrafuncao(x func()) {
	fmt.Println("Prestenção:")
	x()
}

//--------------------------------------------------------------

func Exercicio8() {

	// Demonstre o funcionamento de um closure.
	//- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

	a := tempo()
	b := tempo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func tempo() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}
