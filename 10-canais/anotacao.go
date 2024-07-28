package capitulo10

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Anotacao1() {
	canal := make(chan int)

	go func() {
		canal <- 42
	}()

	fmt.Println(<-canal)

}

func Anotacao2() {

	// Canais direcionais & Utilizando canais

	canal := make(chan int)

	go send(canal)
	receive(canal)

	fmt.Printf("%T", (chan<- int)(canal)) // um exemplo de conversion
}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("O valor recebido, pelo o canal, foi:", <-r)
}

//-----------------------------

func Anotacao3() {

	canal := make(chan int)

	wg.Add(2)

	go loop(10, canal)
	go print(canal)

	wg.Wait()

}

func loop(t int, s chan<- int) {
	defer wg.Done()
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func print(r <-chan int) {
	defer wg.Done()
	for v := range r {
		fmt.Println("recebido:", v)
	}
}

//=============================================

//select

func Anotacao4() {

	//- Exemplo 1:
	//- Duas go funcs enviando X/2 numeros cada uma pra um canal
	//- For loop X valores, select case ←x

	a := make(chan int)
	b := make(chan int)
	x := 500

	wg.Add(2)

	go func(x int) {
		defer wg.Done()
		for i := 0; i < x; i++ {
			a <- i
		}
		close(a)
	}(x / 2)

	go func(x int) {
		defer wg.Done()
		for i := 0; i < x; i++ {
			b <- i
		}
		close(b)
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
	wg.Wait()

}

func Anotacao5() {

	//- Exemplo 2:
	//- Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
	//- Func 2 for infinito, select: case envia pra canal, case recebe de quit

	canal := make(chan int)
	quit := make(chan int)

	wg.Add(2)

	go func1(canal, quit)
	go func2(canal, quit)

	wg.Wait()
}

func func1(canal chan int, quit chan int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func func2(canal chan int, quit chan int) {
	defer wg.Done()
	qualquercoisa := 0
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			close(canal)
			return
		}

	}

}

func Anotacao6() {

	//- Exemplo 3:
	//- Chans par, ímpar, quit
	//- Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
	//- Func receive é um select entre os três canais, encerra no quit
	//- Problema!

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	wg.Add(2)

	go send2(par, impar, quit)
	go receive2(par, impar, quit)

	wg.Wait()
}

func send2(par, impar chan int, quit chan bool) {
	defer wg.Done()
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	quit <- true
	close(par)
	close(impar)
	close(quit)
}

func receive2(par, impar chan int, quit chan bool) {
	defer wg.Done()
	for {
		select {
		case valor := <-par:
			fmt.Println("Valor par: ", valor)
		case valor := <-impar:
			fmt.Println("Valor ímpar: ", valor)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra! Ó:", v)
			} else {
				fmt.Println("Encerrando. Recebemos:", v)
			}
		}
	}
}

//--------------------------------------------------------------------

func Anotacao7() {

	//A expressão comma ok

	canal := make(chan int)

	go func() {
		canal <- 42
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok)

	// Foi adicionado no exercicio anterior o seguinte:
	//
	//case v, ok := <-quit:
	//	if !ok {
	//		fmt.Println("Deu zebra! Ó:", v)
	//	} 	else {
	//			fmt.Println("Encerrando. Recebemos:", v)
	//		}

}

//--------------------------------------------------------------------------

// Convergência
//
// - Observamos convergência quando informação de vários canais é enviada a um número menor de canais.

func Anotacao8() {

	//- Exemplo 1:
	//- Canais par, ímpar, e converge.
	//- Func send manda pares pra um, ímpares pro outro, depois fecha.
	//- Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!

	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	wg.Add(2)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for valor := range converge {
		fmt.Println("Valor recebido:", valor)
	}
}

func envia(par, impar chan int) {
	contador := 100
	for i := 0; i < contador; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func recebe(par, impar, converge chan int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for valor := range par {
			converge <- valor
		}
	}()

	go func() {
		defer wg.Done()
		for valor := range impar {
			converge <- valor
		}
	}()
	wg.Wait()
	close(converge)
}

func Anotacao9() {

	//- 2. Rob Pike (palestra Go Concurrency Patterns):
	//- Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
	//- Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
	//- Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.

	canal := converge(trabalho("Primeiro"), trabalho("Segundo"))

	for x := 0; x < 15; x++ {
		fmt.Println(<-canal)
	}

}

func trabalho(s string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)

	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)

	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}

//---------------------------------------------------------------------

// Divergencia

func Anotacao10() {

	//- Na prática, exemplos:
	//- 1. Um stream vira centenas de go funcs que depois convergem.
	//    - Dois canais.
	//    - Uma func manda X números ao primeiro canal.
	//    - Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
	//    - Trabalho() é um timer aleatório pra simular workload.
	//    - Por fim, range canal dois demonstra os valores.

	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}

}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho2(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho2(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}

func Anotacao11() {

	//- 2. Com throttling! Ou seja, com um número máximo de go funcs.
	//    - Ídem acima, mas a func que lança go funcs é assim:
	//    - Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.

	canal1 := make(chan int)
	canal2 := make(chan int)
	limite := 5

	go manda2(100, canal1)
	go outra2(limite, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}

}

func manda2(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra2(limite int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < limite; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho3(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho3(n int) int {
	time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3)))
	return n
}

//=========================================================================================================

//Context

func Anotacao12() {

	//- Só pra ter uma idéia geral:
	//- Se a gente lança 500 goroutines pra fazer uma tarefa, e cancelamos a tarefa no meio do caminho, como fazemos pra matar as goroutines?
	//- Documentação: https://golang.org/pkg/context/
	//- Aos aventureiros: https://blog.golang.org/context
	//- Destaques:
	//    - ctx := context.Background
	//    - ctx, cancel = context.WithCancel(context.Background)
	//    - goroutine: select case ←ctx.Done(): return; default: continua trabalhando.
	//    - check ctx.Err();
	//    - Tambem tem WithDeadline/Timeout
	//- Exemplos (Todd):
	//    - Analisando:
	//        - Background: https://play.golang.org/p/cByXyrxXUf
	//        - WithCancel: https://play.golang.org/p/XOknf0aSpx
	//        - Função Cancel: https://play.golang.org/p/UzQxxhn_fm
	//    - Exemplos práticos:
	//        - func WithCancel: https://play.golang.org/p/Lmbyn7bO7e
	//        - func WithCancel: https://play.golang.org/p/wvGmvMzIMW
	//        - func WithDeadline: https://play.golang.org/p/Q6mVdQqYTt
	//        - func WithTimeout: https://play.golang.org/p/OuES9sP_yX
	//        - func WithValue: https://play.golang.org/p/8JDCGk1K4P

}
