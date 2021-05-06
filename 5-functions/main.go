package main

import "fmt"

func add(v1 int, v2 int) int {
	return v1 + v2
}

func add2(v1, v2 int) (result int) {
	result = v1 + v2

	return
}

func divide(v1, v2 int) (int, int) {
	return int(v1 / v2), v1 % v2
}

func add3(v1 int) func(int) int {
	return func(v2 int) int {
		return v1 + v2
	}
}

func forEach(items []int, f func(int)) {
	defer fmt.Println("Hello") // defer

	for i := 0; i < len(items); i++ {
		f(items[i])
	}
}

func main() {

	fmt.Println(add3(4)(2))

	items := []int{1, 2, 3}

	forEach(items, func(i int) {
		fmt.Printf("%d * 2 = %d\n", i, i*2)
	})

	_, reminder := divide(5,3)
	fmt.Println(reminder)


}
