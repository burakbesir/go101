package main

import "fmt"

func main() {

	anArr := []int{3, 5, 7, 9}

	for i := 0; i < len(anArr); i++ {
		fmt.Printf("%d ", anArr[i])
	}

	fmt.Println()

	for index, value := range anArr {
		fmt.Printf("Index : %d Value : %d", index, value)
	}

	//While
	var condition bool
	for condition {

	}

	//infinite loop
	//for {
	//
	//}

}
