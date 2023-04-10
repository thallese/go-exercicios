package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		//fmt.Println("Entrou no for", i)
		wg.Add(1)

		numTask := i

		go func() {
			defer wg.Done()
			working(numTask)
		}()
	}

	wg.Wait()
}

func working(id int) {
	fmt.Println("Starting task number: ", id)
	time.Sleep(time.Second)
	fmt.Println("Task number ", id, "done")
}
