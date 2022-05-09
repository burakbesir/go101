package main

import (
	"fmt"
	"time"
)

func sayHello(c chan string) {
	c <- "Hello"
	c <- "World"
}

func main() {

	// Example 1
	//c := make(chan string)
	//go sayHello(c)
	//result := <-c
	//result2 := <-c
	//fmt.Println(result)
	//fmt.Println(result2)

	// Example 2
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(10 * time.Second)

}
