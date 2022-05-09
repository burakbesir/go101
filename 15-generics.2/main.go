package main

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(v T) *Tree[T] {
	return nil
}

var stringTree Tree[string]

func main() {

}
