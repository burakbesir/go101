package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Main start")

	go func() {
		fmt.Println("Hello First Go Routine")
	}()

	go function()

	fmt.Println("Main End")

	time.Sleep(1 * time.Second)
}

func function() {
	fmt.Println("Routine 2")

	go func() {
		fmt.Println("Routine 3")
	}()
}
