package capitulo7

import "fmt"

func Anotacao1() {

	x := 19
	y := &x

	fmt.Println(*y)

	a := 10
	fmt.Println("Sem ponteiro", recebe_valor(a))
	fmt.Println(a)

	fmt.Println("com ponteiro", recebe_ponteiro(&a))
	fmt.Println(a)
}

func recebe_valor(a int) int {
	a++
	return a
}

func recebe_ponteiro(a *int) int {
	*a++
	return *a
}
