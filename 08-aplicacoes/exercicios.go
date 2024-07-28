package capitulo8

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func Exercicio1() {

	//- Partindo do código abaixo, utilize marshal para transformar []user em JSON.

	type user struct {
		First string
		Age   int
	}

	Usuario := []user{{
		First: "James",
		Age:   32,
	}, {
		First: "Moneypenny",
		Age:   27,
	}, {
		First: "M",
		Age:   54,
	}}

	sb, err := json.Marshal(Usuario)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

}

func Exercicio2() {

	// - Partindo do código abaixo, utilize unmarshal e demonstre os valores.

	s := []byte(`[
		{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},
		{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},
		{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}
	]`)
	fmt.Println(s)

	type Personagens struct {
		First   string   `json:"First"`
		Last    string   `json:"Last"`
		Age     int      `json:"Age"`
		Sayings []string `json:"Sayings"`
	}

	var personagens []Personagens
	err := json.Unmarshal(s, &personagens)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(personagens)
}

func Exercicio3() {

	//- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.

	type user struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(users)
}

func Exercicio4() {

	//- Partindo do código abaixo, ordene a []int e a []string.

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

}

func Exercicio5() {

	//- Partindo do código abaixo, ordene os []user por idade e sobrenome.

	type user struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age, v.Sayings)

	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
		// return users[i].Age < users[j].Age
	})
	//sort.Slice(Carros, func(i, j int) bool {
	//  return Carros[i].potencia < Carros[j].potencia
	// })-

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age, v.Sayings)

	}

}
