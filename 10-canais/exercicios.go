package capitulo10

import (
	"fmt"
)

func Exercicio1() {

	//- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
	//- Usando uma função anônima auto-executável
	//- Usando buffer

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// buffer

	d := make(chan int, 1)

	d <- 42

	fmt.Println(<-d)

}

//================================================================================================

func Exercicio2() {

	// Faça esse código funcionar: https://play.golang.org/p/oB-p3KMiH6

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}

//=====================================================================================================

func Exercicio3() {

	//- Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
	//- ...use um for range loop para demonstrar os valores do canal.
	wg.Add(1)

	c := gen()
	go receive3(c)

	fmt.Println("about to exit")
	wg.Wait()
}

func gen() <-chan int {
	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive3(c <-chan int) {
	wg.Add(1)
	defer wg.Done()
	go func() {
		for v := range c {
			fmt.Println("O valor recebido foi:", v)
		}
		wg.Done()
	}()

}

//===============================================================================================

func Exercicio4() {

	//- Utilizando este código: https://play.golang.org/p/MvL6uamrJP
	//- ...use um select statement para demonstrar os valores do canal.
	wg.Add(1)

	c := gen2()
	go receive4(c)

	fmt.Println("about to exit")
	wg.Wait()
}

func gen2() <-chan int {
	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive4(c <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-c:
			for v := range c {
				fmt.Println("Valor recebido:", v)
			}
		case _, ok := <-c:
			if !ok {
				fmt.Println("Encerrado:")
				return
			}
		}
	}
}

//==========================================================================

func Exercicio5() {

	//- Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
	//- ...demonstre o comma ok idiom.

	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

}

//=====================================

func Exercicio6() {

	//- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

	canal := make(chan int)

	wg.Add(2)

	go coloca(canal)
	go mostra(canal)

	wg.Wait()

}

func coloca(c chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func mostra(c chan int) {
	defer wg.Done()
	for valor := range c {
		fmt.Println("Valor:", valor)
	}

}

//===============================================

func Exercicio7() {

	//- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
	//- Tire estes números do canal e demonstre-os.

	// - A forma que estou fazendo não é convergencia, mas quis colocar esse nome
	wg.Add(2)

	canal := make(chan int)

	go lanca(canal)
	go demonstrar(canal)

	wg.Wait()

	// Não consegui fazer da maneira que eu queria
}

func lanca(canal chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				canal <- i
			}
		}()
	}
	wg.Wait()
	close(canal)
}

func demonstrar(canal chan int) {
	for {
		select {
		case _, ok := <-canal:
			if !ok {
				fmt.Println("Encerrando..")
				return
			}
		}
		for valor := range canal {
			fmt.Println("Valor recebido:", valor)
		}
	}
}

func CorrecaoExercicio7() {
	canal := make(chan int)

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}
	for valor := range canal {
		fmt.Println("Valor recebido:", valor)
	}

	wg.Wait()
}
