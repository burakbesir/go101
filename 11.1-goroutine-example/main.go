package main

import (
	"fmt"
	"time"
)

func main() {

	//Example 1
	//for i := 1; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	time.Sleep(1 * time.Second)

	////Example 2
	//for i := 1; i < 10; i++ {
	//	go print(i)
	//}
	//time.Sleep(1 * time.Second)

	//// Example 3
	for i := 1; i < 10; i++ {
		go print2(&i)
	}
	time.Sleep(1 * time.Second)
}

func print(i int) {
	fmt.Println(i)
}

func print2(i *int) {
	fmt.Println(*i)
}
