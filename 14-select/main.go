package main

import (
	"fmt"
	"time"
)

func main() {

	numbers := make(chan int)
	strings := make(chan string)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			strings <- fmt.Sprintf("A%d", i)
		}
	}()

	go func() {
		sumOfNumbers := 0
		stringCounter := 0
		for {
			select {
			case number := <-numbers:
				sumOfNumbers += number
			case <-strings:
				stringCounter++
			case <-quit:
				fmt.Printf("Sum : %d , Count : %d", sumOfNumbers, stringCounter)
				break
			}
		}
	}()

	time.Sleep(1 * time.Second)
	quit <- 0

	time.Sleep(1 * time.Second)

}
