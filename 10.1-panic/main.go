package main

import "fmt"

func main() {

	//Example 1
	//fmt.Println("11111")
	//panic("Hello")
	//fmt.Println("22222")

	// Example 2

	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	fmt.Println("11111")
	panic("Hello")
	fmt.Println("22222")
}
