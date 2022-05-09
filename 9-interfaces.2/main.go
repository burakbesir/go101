package main

type Order struct {
}

type Repository interface {
	Use()
}

type OrderRepository interface {
	getById(int) Order
}

type test struct {
	Repository
}

func (o *test) getById(id int) Order {
	return Order{}
}

func main() {

	var r Repository

	r = test{}

	r.Use()

}
