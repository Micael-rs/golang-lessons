package capitulo5

import (
	"fmt"
)

func Exercicio1() {

	//- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
	//- Nome
	//- Sobrenome
	//- Sabores favoritos de sorvete
	//- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

	type Pessoa struct {
		nome      string
		sobrenome string
		gostos    []string
	}

	pessoa1 := Pessoa{
		nome:      "Micael",
		sobrenome: "Ramos,",
		gostos:    []string{"Chocolate", "Baunilha", "Morango"},
	}

	fmt.Println(pessoa1)
}

func Exercicio2() {

	//- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
	//- Demonstre os valores do map utilizando range.
	//- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

	type Pessoa struct {
		nome      string
		sobrenome string
		gostos    []string
	}

	usuario := make(map[string]Pessoa)

	usuario["Ramos"] = Pessoa{
		nome:      "Micael",
		sobrenome: "Ramos",
		gostos:    []string{"Chocolate", "Morango", "Pixi"},
	}

	usuario["Santana"] = Pessoa{"Isabelli", "Santana", []string{"Micael", "Micael", "Micael"}}

	for chave, valor := range usuario {
		fmt.Println(chave)
		for _, v := range valor.gostos {
			fmt.Print(v, " ")
		}
		fmt.Println(" ")
	}
}

func Exercicio3() {

	//- Crie um novo tipo: veículo
	//- O tipo subjacente deve ser struct
	//- Deve conter os campos: portas, cor
	//- Crie dois novos tipos: caminhonete e sedan
	//- Os tipos subjacentes devem ser struct
	//- Ambos devem conter "veículo" como struct embutido
	//- O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
	//- O tipo sedan deve conter um campo bool chamado "modeloLuxo"
	//- Usando os structs veículo, caminhonete e sedan:
	//- Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
	//- Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
	//- Demonstre estes valores.
	//- Demonstre um único campo de cada um dos dois.

	type Veiculo struct {
		porta int
		cor   string
	}

	type Caminhonete struct {
		Veiculo
		tracao_nas_quatro bool
		modelo_luxo       bool
	}

	type Sedan struct {
		Veiculo
		tracao_nas_quatro bool
		modelo_luxo       bool
	}

	fiat_strada := Caminhonete{
		Veiculo: Veiculo{
			porta: 4,
			cor:   "Vermelha",
		},
		tracao_nas_quatro: false,
		modelo_luxo:       false,
	}

	toyota_corolla := Sedan{
		Veiculo: Veiculo{
			porta: 4,
			cor:   "Cinza",
		},
		tracao_nas_quatro: false,
		modelo_luxo:       false,
	}

	fmt.Println(fiat_strada)
	fmt.Println(toyota_corolla)

	fmt.Println(fiat_strada.Veiculo)
	fmt.Println(toyota_corolla.Veiculo)
}

func Exercicio4() {

	//- Crie e use um struct anônimo.
	//- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

	Cliente := struct {
		nome       string
		veiculos   map[string]string
		familiares []string
	}{
		nome: "Fernando",
		veiculos: map[string]string{
			"carro do tio":      "sedan",
			"carro do sobrinho": "fusca",
		},
		familiares: []string{"esposa", "filho", "Mãe"},
	}

	fmt.Print(Cliente)
}
