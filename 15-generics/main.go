package main

import "fmt"

func main() {

	m1 := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(sumOfMapValues[string, int64](m1))

	m2 := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
	}
	fmt.Println(sumOfMapValues(m2))

	forInt := sumOfMapValues[string, int64]

	fmt.Println(forInt(m1))

}

func sumOfMapValues[K comparable, V int64 | float64](m map[K]V) V {
	var result V
	for _, v := range m {
		result += v
	}
	return result
}

type Number interface {
	int64 | float64
}
