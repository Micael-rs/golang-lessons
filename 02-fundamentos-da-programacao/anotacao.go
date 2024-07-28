package capitulo2

import(
	"fmt"
	"runtime"
)

func Anotacao(){

//- Sempre que você ver operadores relacionais, o resultado da expressão será um valor booleano.
//- Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.
//- Na prática:
//    - Zero value
//    - Atribuindo um valor
//    - Bool como resultado de operadores relacionais

x := true

fmt.Println(x) 

x = (10>100)
y := (10 == 100)
z := (10<100)

fmt.Println(x, y, z)

//convertendo em rune e em bytes

a := "oi"
b := "é"

c := []byte(a)
d := []rune(b)

fmt.Printf("%v, %v\n", c, d)

//Descobrindo o tipo do meu pc

fmt.Println(runtime.GOOS)
fmt.Println(runtime.GOARCH)

//bimario, hexadecimal e decimal
// base-10: decimal, 0-9
// base-2: binario, 0-1
// base-16: hexadecimal, 0-f

valor := 100

fmt.Printf("%d\t %b\t %#x", valor, valor, valor)

//const

const iota = iota
fmt.Println(iota)

//deslocamento de binario

valor = 10
valor1 := valor>>1

fmt.Printf("%b", valor1)

}

