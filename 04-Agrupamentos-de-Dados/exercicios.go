package capitulo4

import (
	"fmt"
)

func Exercicio1() {

	//- Usando uma literal composta:
	//- Crie um array que suporte 5 valores to tipo int
	//- Atribua valores aos seus índices
	//- Utilize range e demonstre os valores do array.
	//- Utilizando format printing, demonstre o tipo do array.

	array := [5]int{1, 2, 3, 4, 5}

	for _, valor := range array {
		fmt.Print(valor, " ")
	}

	fmt.Printf("%T", array)

}

func Exercicio2() {

	//- Usando uma literal composta:
	//- Crie uma slice de tipo int que suporte 5 valores to tipo int
	//- Atribua valores aos seus índices
	//- Utilize range e demonstre os valores do array.
	//- Utilizando format printing, demonstre o tipo do array.

	slice := []int{1, 2, 3, 4, 5}

	for _, valor := range slice {
		fmt.Print(valor, " ")
	}

	fmt.Printf("%T", slice)

}

func Exercicio3() {

	//- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	//- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

	numeros := []int{1, 2, 3, 4, 5, 6, 7}

	a := numeros[2:6]

	fmt.Println(a)
	fmt.Println(a[2 : len(numeros)-1])

}

func Exercicio4() {

	//- Começando com a seguinte slice:
	//- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//- Anexe a ela o valor 52;
	//- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
	//- Demonstre a slice;
	//- Anexe a ela a seguinte slice:
	//- y := []int{56, 57, 58, 59, 60}
	//- Demonstre a slice x.

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)

	fmt.Println(x)

}

func Exercicio5() {

	//- Comece com essa slice:
	//- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//- Utilizando slicing e append, crie uma slice y que contenha os valores:
	//- [42, 43, 44, 48, 49, 50, 51]

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)

	fmt.Println(y)
	fmt.Println(x)

	z := &x

	fmt.Println(*z)
}

func Exercicio6() {

	//- Crie uma slice usando make que possa conter todos os estados do Brasil.
	// - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
	//- Demonstre o len e cap da slice.
	//- Demonstre todos os valores da slice sem utilizar range.

	estados := make([]string, 26)

	nome := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	for contador := 0; contador < len(estados); contador++ {
		estados[contador] = nome[contador]
	}

	fmt.Println(len(estados), cap(estados))

	for contador := 0; contador < len(estados); contador++ {
		fmt.Println(estados[contador])
	}
}

func Exercicio7() {

	//- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
	//- "Nome", "Sobrenome", "Hobby favorito"
	//- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

	dados := [][]string{
		{
			"Micael",
			"Ramos",
			"Volta dos que não foram",
		},
		{
			"Isabelli",
			"Santana",
			"Avatar",
		},
	}

	for contador := 0; contador < len(dados); contador++ {
		fmt.Println(dados[contador])
	}
}

func Exercicio8() {

	//- Crie um map com key tipo string e value tipo []string.
	//- Key deve conter nomes no formato sobrenome_nome
	//- Value deve conter os hobbies favoritos da pessoa
	//- Demonstre todos esses valores e seus índices.

	pessoa := map[string][]string{
		"Micael_Ramos": {
			"correr", "nadar", "musculação",
		},
		"Isabelli_Santana": {
			"Ler", "Estudar", "DORMIR",
		},
	}

	for chave, valor := range pessoa {
		fmt.Println(chave)
		for _, valor := range valor {
			fmt.Println(valor)

		}
	}
}

func Exercicio9() {

	//- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

	pessoa := map[string][]string{
		"Micael_Ramos": {
			"correr", "nadar", "musculação",
		},
		"Isabelli_Santana": {
			"Ler", "Estudar", "DORMIR",
		},
	}

	pessoa["Daniel"] = []string{
		"Estudar", "dormir", "dormir",
	}

	for chave, valor := range pessoa {
		fmt.Println(chave)
		for _, valor := range valor {
			fmt.Println(valor)

		}
	}

}

func Exercicio10() {

	//- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.

	pessoa := map[string][]string{
		"Micael_Ramos": {
			"correr", "nadar", "musculação",
		},
		"Isabelli_Santana": {
			"Ler", "Estudar", "DORMIR",
		},
	}

	pessoa["Daniel"] = []string{
		"Estudar", "dormir", "dormir",
	}

	delete(pessoa, "Isabelli_Santana")

	for chave, valor := range pessoa {
		fmt.Println(chave)
		for _, valor := range valor {
			fmt.Println(valor)
		}
	}

}
