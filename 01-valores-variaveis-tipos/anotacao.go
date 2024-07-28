package capitulo1

import(
	"fmt"
)

//Quando a variavel é declarada fora da code-block ela esta visivel para os outros pacotes.
//lembrando que fora da code-block deve se usar o VAR, não deve atribuir valor a ela
var a int 

func Anotacao(){

	a = 10
	fmt.Println(a)

	//Dentro da code-block é melhor usar :=, pois declara e atribui valor a ela
	b := 20
	fmt.Println(b)

	//Criando o nosso typo (não é importante por agora)
	type exemplo int
	var c exemplo
	fmt.Println(c)

	//converter o typo exemplo para int
	d := 0
	d = int(c)
	fmt.Println(d)

}