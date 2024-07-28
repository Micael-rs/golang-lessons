package capitulo3

import (
	"fmt"
)

func Anotacao() {

	//- For: Inicialização, condição, pós

	for a := 0; a < 10; a++ {

		fmt.Println(a)

	}
}

func Repeticaoo_hierarquica() {

	//Repetição hierárquica

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Horas: ", horas)
		for contador := 0; contador < 60; contador++ {
			fmt.Print(contador, " ")
		}
		fmt.Println("")
	}

	for ano := 0; ano <= 2; ano++ {
		fmt.Println("Ano ", ano)
		for mes := 0; mes <= 12; mes++ {
			fmt.Println("Mes: ", mes)
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(dia, " ")
			}
			fmt.Println(" ")
		}
		fmt.Println(" ")
	}

}

func Break_continue() {

	//Breack no loop

	contador := 0

	for {

		if contador <= 10 {
			fmt.Println(contador)
			contador++
		} else {
			break
		}

	}

	for contador = 0; contador <= 10; contador++ {
		if contador%2 != 0 {
			continue
		}
		fmt.Println(contador)
	}
}

func Switch() {

	x := 10

	switch {
	case x < 5:
		fmt.Println("mano")
	case x == 10:
		fmt.Println("Brabo")
		//Adicionando fallthrough
		fallthrough
	case x > 5:
		fmt.Println("xiiii")
	}

	y := 5

	switch 5 {
	case y:
		fmt.Println("certo")
	case 10:
		fmt.Println("errado")

	}

	switch y {
	case 0, 2:
		fmt.Println("Entre 0 a 2")
	case 3, 5:
		fmt.Println("Entre 3 a 6")
	default:
		fmt.Println("Entre 7 adiante")
	}

	//&&(e), ||(ou)

}
