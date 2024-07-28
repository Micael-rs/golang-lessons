package capitulo1

import "fmt"

func Exercicio1() {

	//Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
	//1. 42
	//2. "James Bond"
	//3. true
	//- Agora demonstre os valores nestas variáveis, com:
	//1. Uma única declaração print
	//2. Múltiplas declarações print

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

var x int
var y string
var z bool
func Exercicio2(){

	//- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
	//1. Identificador "x" deverá ter tipo int
	//2. Identificador "y" deverá ter tipo string
	//3. Identificador "z" deverá ter tipo bool
	//- Na função main:
	//1. Demonstre os valores de cada identificador
	//2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?

	fmt.Println(x, y, z)
	fmt.Printf("x: %v %T, y: %v %T, z: %v %T", x, x, y, y, z, z)

}

func Exercicio3(){

	//- Utilizando a solução do exercício anterior:
    //1. Em package-level scope, atribua os seguintes valores às variáveis:
	//1. para "x" atribua 42
    //2. para "y" atribua "James Bond"
    //3. para "z" atribua true
    //2. Na função main:
    //1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
    //2. Demonstre a variável "s".

	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprint(x, y, z)
	fmt.Println(s)

}

func Exercicio4(){

	//- Crie um tipo. O tipo subjacente deve ser int.
	//- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
	//- Na função main:
	//	1. Demonstre o valor da variável "x"
	//	2. Demonstre o tipo da variável "x"
	//	3. Atribua 42 à variável "x" utilizando o operador "="
	//	4. Demonstre o valor da variável "x"

	type exemplo int
	var x exemplo

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%v", x)

}

func Exercicio5(){

	//- Utilizando a solução do exercício anterior:
    //1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    //2. Na função main:
    //    1. Isto já deve estar feito:
    //        1. Demonstre o valor da variável "x"
    //        2. Demonstre o tipo da variável "x"
    //        3. Atribua 42 à variável "x" utilizando o operador "="
    //        4. Demonstre o valor da variável "x"
    //    2. Agora faça tambem:
    //        1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
    //        2. Demonstre o valor de "y"
    //        3. Demonstre o tipo de "y"

	type exemplo int
	var x exemplo

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%v\n", x)

	var y int 
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}