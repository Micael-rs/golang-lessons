package capitulo6

import (
	"fmt"
	"math"
)

// video 1 e video 2

func Anotacao1() {

	resultado := Soma(10, 10)
	fmt.Println(resultado)

	Bomdia()

	total := Sum(10, 10, 30, 20)
	fmt.Println(total)

	s := []int{10, 10, 30, 20}
	a := Sum(s...)

	fmt.Println(a)
}

func Soma(a, b int) int {
	return a + b
}

func Bomdia() {
	fmt.Println("bom dia")
}

func Sum(a ...int) int {
	total := 0

	for _, a := range a {
		total += a
	}
	return total
}

// video 3

func Anotacao2() {

	defer fmt.Print("b")
	fmt.Println("a")

}

// video 4: Métodos

type pessoa struct {
	nome  string
	idade int
}

func Anotacao3() {

	cliente1 := pessoa{"Micael", 18}

	cliente1.Indentificador()

}

func (p pessoa) Indentificador() {

	fmt.Println(p.nome, "é o cliente")

}

// video 5:  Interfaces & polimorfismo

//- Em Go, valores podem ter mais que um tipo.
//- Uma interface permite que um valor tenha mais que um tipo.
//- Declaração: keyword identifier type → type x interface
//- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
//- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
//- Esse tipo será o seu tipo e também o tipo da interface.

//- Exemplos:
//- Os tipos profissão1 e profissão2 contem o tipo pessoa
//- Cada um tem seu método oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()
//- Implementam a interface gente
//- Ambos podem acessar o função serhumano() que chama o método oibomdia() de cada gente
//- Tambem podemos no método serhumano() tomar ações diferentes dependendo do tipo:
//    switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*

type pedreiro struct {
	pessoa
	salario float64
}

type bananeiro struct {
	pessoa
	salario float64
}

func (a pedreiro) Oibomdia() {
	fmt.Println(a.nome, ", tem um bom dia, com uma obra bem feita!")
}

func (a bananeiro) Oibomdia() {
	fmt.Println(a.nome, ", tem um bom dia, com muitas demandas!")
}

type gente interface {
	Oibomdia()
}

func Serhumano(a gente) {

	a.Oibomdia()

	switch a := a.(type) {
	case pedreiro:
		fmt.Println(a.nome, "Recebe:", a.salario)
	case bananeiro:
		fmt.Println(a.nome, "Recebe:", a.salario)
	}

}

func Anotacao4() {

	pedreiro := pedreiro{
		pessoa:  pessoa{"Carlinhos", 40},
		salario: 14000.06,
	}

	bananeiro := bananeiro{
		pessoa:  pessoa{"Elias", 40},
		salario: 30000,
	}

	//pedreiro.Oibomdia()
	//Serhumano(pedreiro)
	//fmt.Println(" ")
	//bananeiro.Oibomdia()
	//Serhumano(bananeiro)
	//fmt.Println(" ")

	// usando a interface

	Serhumano(pedreiro)
	fmt.Println(" ")
	Serhumano(bananeiro)

}

// continuação - video 5

type circulo struct {
	raio float64
}

type triangulo struct {
	largura, altura float64
}

func (c circulo) Area() float64 {
	return math.Pi * c.raio * c.raio
}

func (t triangulo) Area() float64 {
	return t.altura * t.largura
}

type soma interface {
	Area() float64
}

func Mostrar_area(soma soma) {
	fmt.Println(soma.Area())
}

func Anotacao5() {

	c := circulo{raio: 5}
	t := triangulo{largura: 4, altura: 3}

	Mostrar_area(c)
	Mostrar_area(t)

}

// video 6 e 7: função anônima e expressão

func Anotacao6() {

	x := 10

	func(x int) {
		fmt.Println(x, "Vezes 20")
		fmt.Println(x * 20)
	}(x)

	y := (x)

	fmt.Println(y)

	b := 0

	a := func(b int) int {
		return b * 10
	}

	fmt.Println(a(b))

}

// video 8: retornando função

func Anotacao7() {

	fmt.Println(Retorna_função()(3))

}
func Retorna_função() func(int) int {
	return func(i int) int {
		return i * 10
	}
}

// video 9: Callback

func Anotacao8() {

	funcao := Divisao(Adicao, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)

	fmt.Println(funcao)
}

func Adicao(x ...int) int {
	a := 0

	for _, v := range x {
		a = a + v
	}

	return a
}

func Divisao(f func(x ...int) int, y ...int) int {
	var lista []int
	for _, v := range y {
		if v%2 != 0 {
			lista = append(lista, v)
		}
	}

	total := f(lista...)
	return total
}

// video 10: Closure

func Anotacao9() {

	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//  video 11: Recursividade

func Anotacao10() {
	fmt.Println(fatoriais(4))
}

func fatoriais(x int) int {
	if x == 1 {
		return x
	}
	return x * fatoriais(x-1)
}
