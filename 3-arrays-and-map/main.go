package main

import "fmt"

func main() {

	var names [3]string

	// slice
	anArr := []int{1, 2, 3, 4}

	fmt.Println(names, anArr)

	anArr = append(anArr, 34)

	anotherArr := make([]int, 0, 5)

	fmt.Println(anotherArr)

	// maps

	m := make(map[string]string)

	m["Hello"] = "World"

	fmt.Println(m["Hello"])

	v, e := m["Hello"]
	fmt.Println(v, e)

	if e {
		fmt.Println(v)
	}

	// if
	if v, e := m["Hello"]; e {
		fmt.Println(v)
	}

	switch m["Hello"] {
	case "World":
	case "Another":
	default:

	}

}
