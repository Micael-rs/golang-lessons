package capitulo8

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

// video 1: jason marshal

func Anotacao1() {

	// Um código em GO será passado para json

	type Cliente struct {
		Nome   string
		Email  string
		Numero int
		Vicios bool
	}

	a := &Cliente{
		Nome:   "Micael",
		Email:  "micael@gmail",
		Numero: 11999178730,
		Vicios: false,
	}

	b, err := json.Marshal(*a)
	if err != nil {
		fmt.Println("erro", err)
	}

	os.Stdout.Write(b)

}

// video 2: unmarshal

func Anotacao2() {

	var Jason = []byte(`{"Nome":"Micael","Email":"micael@gmail","Numero":11999178730,"Vicios":false}`)

	type Informacoes struct {
		Nome   string `json:"Nome"`
		Email  string `json:"Email"`
		Numero int    `json:"Numero"`
		Vicios bool   `json:"Vicios"`
	}

	var info Informacoes
	err := json.Unmarshal(Jason, &info)
	if err != nil {
		fmt.Println("erro", err)
	}

	fmt.Println(info)
}

// continuação.. video 2: com enconde (Utiliza interface)

func Anotacao3() {

	type Cliente struct {
		Nome   string
		Email  string
		Numero int
		Vicios bool
	}

	a := &Cliente{
		Nome:   "Micael",
		Email:  "micael@gmail",
		Numero: 11999178730,
		Vicios: false,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(a)
}

// video 3: sort

func Anotacao4() {

	a := []int{12, 37, 2, 71, 50, 37}

	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

	//sort.Strings()
}

// Criando a nossa Sort

type Carro struct {
	Nome              string
	consumo, potencia int
}

type OrdenarPorPotencia []Carro

func (x OrdenarPorPotencia) Len() int { return len(x) }

func (x OrdenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }

func (x OrdenarPorPotencia) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func Anotacao5() {

	//- Exemplo:
	//- struct carros: nome, consumo, potencia
	//- slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
	//- tipo ordenarPorPotencia
	//- tipo ordenarPorConsumo
	//- tipo ordenarPorLucroProDonoDoPosto

	Carro := &[]Carro{
		{"Camaro", 100, 230},
		{"Chevete", 40, 30},
		{"Fusca", 50, 39},
	}

	fmt.Println("Antes de ordenar:", Carro)
	sort.Sort(OrdenarPorPotencia(*Carro))
	fmt.Println("Depois de ordenar:", Carro)

	//sort.Slice(Carros, func(i, j int) bool {
	//  return Carros[i].potencia < Carros[j].potencia
	// })
}

// bcrypt
func Anotacao6() {

	// - É uma maneira de encriptar senhas utilizando hashes.

	senha := "22Junho2002"
	senhaerrada := "abcde"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))
}
