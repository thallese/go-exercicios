package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//var mu sync.Mutex
	totalGoroutines := 1000
	wg.Add(totalGoroutines)
	j := 0

	for i := 1; i <= totalGoroutines; i++ {
		go func() {
			//mu.Lock()
			j++
			runtime.Gosched()
			//mu.Unlock()
		}()
		wg.Done()
	}

	fmt.Println("Total de go routines:", totalGoroutines)
	fmt.Println("Total do contador: ", j)
	fmt.Println("Foram perdidas", totalGoroutines-j, "incrementações")
	wg.Wait()

}
