package capitulo4

import (
	"fmt"
)

func Anotacao() {

	//declarando array
	a := []int{1, 0, 2, 3, 5, 6}
	//declarando uma slice
	var c []int
	//corte
	b := a[1:4]

	fmt.Println(a, b, c)
	//mostra quantos intens há na array
	fmt.Println(len(a))

	//adicionando intem na slice
	c2 := append(c, 2)

	fmt.Println(c2)

}

func Anotacao2() {

	frutas := []string{"banana", "Maçã", "Melancia"}

	for id, nome := range frutas {
		fmt.Println("ID:", id, "Nome:", nome)
	}

	//caso eu fosse adicionar algo
	frutas = append(frutas, "Laranja")

	for id, nome := range frutas {
		fmt.Println("ID:", id, "Nome:", nome)
	}

}

func Anotacao3() {

	slices := []int{5, 5, 5, 5, 5}
	total := 0

	for _, valor := range slices {
		total += valor
	}

	fmt.Print("O valor total dessa slice é: ", total)

}

func Anotacao4() {

	mes := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho"}

	a := mes[:]

	fmt.Println(a)

	for i := 0; i < len(mes); i++ {
		fmt.Println(mes[i])
	}
}

func Anotacao5() {

	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	c := append(a, b...)
	//ou a := append(a, b...)

	fmt.Println(c)

	//apagando conteudo na slice
	c = append(c[:2], c[3:]...)

	fmt.Println(c)
}

func Anotacao6() {

	slices := make([]int, 5, 10)

	for i := 0; i < 5; i++ {
		slices[i] = i + 1
	}

	fmt.Println(slices, len(slices), cap(slices))

	slices = append(slices, 5, 6, 7, 8)

	fmt.Println(slices, len(slices), cap(slices))

}

func Anotacao7() {

	a := [][]int{
		{1, 2, 3}, //0
		{4, 5, 6}, //1
		{7, 8, 9}, //2
	}

	//Caso eu queira selecionar a linha 1 e vertical 1, que seria cinco, faça:
	//a[1][1]

	fmt.Println(a[0])
	fmt.Println(a[1][1])

	b := [][][]int{
		{
			{1, 2, 3, 4},
		},
	}

	fmt.Println(b[0][0][0])
}

func Anotacao8() {

	//Criando um map
	contatos := map[string]int{
		"joana": 12345,
	}

	fmt.Println(contatos)

	//adicionando outro valor
	contatos["micael"] = 11999178730

	fmt.Println(contatos["micael"])

	//exemplo do tour golang
	type Localizacao struct {
		lat, long float64
	}

	a := map[string]Localizacao{
		"Rio de janeiro": {
			1111, 1111,
		},
	}

	fmt.Println(a)

	a["São paulo"] = Localizacao{
		2222, 3333,
	}

	fmt.Println(a["São paulo"])
	fmt.Println(a["Rio de janeiro"])

	//validando uma informação dentro da map a
	//comma on idiom

	if _, ok := a["africa"]; !ok {
		fmt.Println("Não achado")
	} else {
		fmt.Println("Encontrado")
	}

	validacao := "mexico"

	if _, ok := a[validacao]; !ok {
		fmt.Println("Não achado")
	} else {
		fmt.Println("Encontrado")
	}

	//deletando aldo da map

	delete(a, "São paulo")

	fmt.Println(a)
}
