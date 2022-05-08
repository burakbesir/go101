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
		Line: []Line{
			Line{Id: 14},
			{Id: 13}, // if declaration has multiple line last line should end with a comma
		},
	}

	fmt.Println(order.Id)
}
