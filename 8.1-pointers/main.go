package main

import "fmt"

type Human struct {
	Age int
}

func (h Human) happyBirthDay() {
	h.Age++
}

func (h *Human) happyBirthDay2() {
	h.Age++
}

func main() {

	h := Human{Age: 5}

	h.happyBirthDay()
	fmt.Println(h.Age)

	h.happyBirthDay2()
	fmt.Println(h.Age)
}
