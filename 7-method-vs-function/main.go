package main

import "fmt"

type Human struct {
	Name string
}

//method
func (h Human) sayHello() {
	fmt.Printf("Hello %s", h.Name)
}

//function
func functionSayHello() {
	fmt.Println("Hello")
}

func main() {

	functionSayHello()

	human := Human{Name: "Rahman"}

	Human.sayHello(human)
	human.sayHello()

}
