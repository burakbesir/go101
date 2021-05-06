package main

import "fmt"

type Order struct {
	Id     int
	Number string
	Line   []Line
}

type Line struct {
	Id int
}

func main() {

	order := Order{
		Id:     12,
		Number: "13",
		Line:   []Line{{Id: 14}},
	}

	fmt.Println(order.Id)
}
