package capitulo9

import (
	"fmt"
	"runtime"
)

func Exercicio1() {

	wg.Add(2)

	go mensagem1()
	go mensagem2()

	wg.Wait()
}

func mensagem1() {
	fmt.Println("Olá")
	wg.Done()
}

func mensagem2() {
	fmt.Println("Mundo")
	wg.Done()
}

//----------------------------------------------------

//- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//- Crie um tipo para um struct chamado "pessoa"
//- Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//- Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//- Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//- Demonstre no seu código:
// - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//- Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p *pessoa) falar() {
	fmt.Println("Prazer, me chamo", p.nome, p.sobrenome, "e tenho", p.idade)
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func Exercicio2() {

	pessoa := pessoa{"Micael", "Ramos", 18}

	dizerAlgumaCoisa(&pessoa)
}

//=========================================================================

//- Utilizando goroutines, crie um programa incrementador:
//- Tenha uma variável com o valor da contagem
//- Crie um monte de goroutines, onde cada uma deve:
//- Ler o valor do contador
//- Salvar este valor
//- Fazer yield da thread com runtime.Gosched()
//- Incrementar o valor salvo
//- Copiar o novo valor para a variável inicial
//- Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//- Demonstre que há uma condição de corrida utilizando a flag -race

func Exercicio3() {

	Goroutines := 500
	contador := 0

	wg.Add(Goroutines)

	for i := 0; i < Goroutines; i++ {
		go func() {
			mu.Lock()
			runtime.Gosched()
			contador++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
