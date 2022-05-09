package main

import "fmt"

type Order struct {
	id     int
	Number string
	Line   []Line
}

//type ID int

type Line struct {
	Id int
}

func main() {

	order := Order{
		id:     12,
		Number: "13",
		Line: []Line{
			Line{Id: 14},
			{Id: 13}, // if declaration has multiple line last line should end with a comma
		},
	}

	fmt.Println(order.id)
}
