package capitulo5

import (
	"fmt"
)

func Anotacao1() {

	type Cliente struct {
		nome   string
		email  string
		numero int
		vicios bool
	}

	a := &Cliente{
		"Micael",
		"micael@gmail",
		11999178730,
		false,
	}

	fmt.Println(*a)
}

func Anotacao2() {

	type Pessoa struct {
		nome  string
		idade int
	}

	type Trabalho struct {
		Pessoa
		cargo string
	}

	funcionario := Trabalho{
		Pessoa: Pessoa{"Micael", 18},
		cargo:  "Programador",
	}

	fmt.Println(funcionario)
	fmt.Println(funcionario.nome)
	fmt.Println(Trabalho{})
}

func Anotacao3() {

	a := struct {
		nome  string
		idade int
	}{
		nome:  "micael",
		idade: 10,
	}

	fmt.Println(a)
}
