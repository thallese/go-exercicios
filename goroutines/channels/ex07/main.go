package main

import "fmt"

func main() {
	c := make(chan int)
	colocaNoCanal(c)
	retiraDoCanal(c)
}

func colocaNoCanal(c chan int) {

	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()

}

func retiraDoCanal(c chan int) {
	for v := range c {
		n, err := fmt.Println(v)
		fmt.Println(n)
		fmt.Println(err)
	}
}
