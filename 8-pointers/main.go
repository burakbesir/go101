package main

import "fmt"

type Human struct {
	Age int
}

func (h Human) printAge() {
	fmt.Printf("Age : %d\n", h.Age)
}

func changeMyAge(h Human) {
	h.Age = 12
	fmt.Printf("Set Age for Human %v\n", h)
}

func changeMyAge2(h *Human) {
	h.Age = 12
	fmt.Printf("Set Age for Human %v\n", h)
}

func main() {

	h := Human{Age: 5}

	changeMyAge(h)
	h.printAge()

	changeMyAge2(&h)
	h.printAge()
}
