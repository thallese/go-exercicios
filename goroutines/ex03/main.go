/*
- Utilizando goroutines, crie um programa incrementador:
  - Tenha uma variável com o valor da contagem
  - Crie um monte de goroutines, onde cada uma deve:
  - Ler o valor do contador
  - Salvar este valor
  - Fazer yield da thread com runtime.Gosched()
  - Incrementar o valor salvo
  - Copiar o novo valor para a variável inicial
  - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
  - Demonstre que há uma condição de corrida utilizando a flag -race
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	totalGoroutines := 100000000
	wg.Add(totalGoroutines)
	j := 0

	for i := 1; i <= totalGoroutines; i++ {
		go func() {
			j++
			runtime.Gosched()
		}()
		wg.Done()
	}

	fmt.Println("Total de go routines:", totalGoroutines)
	fmt.Println("Total do contador: ", j)
	fmt.Println("Foram perdidas", totalGoroutines-j, "incrementações")
	wg.Wait()

}
