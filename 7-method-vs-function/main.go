package main

import "fmt"

//function
func functionSayHello() {
	fmt.Println("Hello")
}

type Human struct {
	Name string
}

//method
func (h Human) sayHello() {
	fmt.Printf("Hello %s", h.Name)
}

func main() {

	functionSayHello()

	human := Human{Name: "Rahman"}

	Human.sayHello(human)
	human.sayHello()

}
