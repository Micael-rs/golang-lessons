package capitulo9

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// - Goroutines & WaitGroups

//  Concorrência vs. paralelismo
//- Goroutines!
//- O que são goroutines? São "threads."
//- O que são threads? [WP](https://pt.wikipedia.org/wiki/Thread_...)
//- Na prática: go func.
//- Exemplo: código termina antes da go func executar.
//- Ou seja, precisamos de uma maneira pra "sincronizar" isso.
//- Ah, mas então... não.
//- Qualé então? sync.WaitGroup:
//- Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
//- func Add: "Quantas goroutines?"
//- func Done: "Deu!"
//- func Wait: "Espera todo mundo terminar."
//- Ah, mas então... sim!
//- Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()

var wg sync.WaitGroup
var mu sync.Mutex

func Anotacao1() {

	wg.Add(2)

	go funcao1()
	go funcao2()

	wg.Wait()
}

func funcao1() {
	for i := 0; i < 100; i++ {
		fmt.Println("função 1:", i)
		time.Sleep(2000)
	}
	wg.Done()
}

func funcao2() {
	for i := 0; i < 100; i++ {
		fmt.Println("função 2:", i)
	}
	wg.Done()

}

//-----------------------------------------------

func Anotacao2() {

	// Utilizando Mutex
	// - exemplo simples de como evitar Condição de Corrida
	// - go run -race main.go

	contador := 0
	goroutines := 1000

	wg.Add(goroutines)

	fmt.Println("CPUs:", runtime.NumCPU())

	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor Final:", contador)
}

//--------------------

func Anotacao3() {

	var contador int64
	contador = 0
	goroutines := 1000

	wg.Add(goroutines)

	fmt.Println("CPUs:", runtime.NumCPU())

	for i := 0; i < goroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Valor Final:", contador)

}
