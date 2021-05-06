package main

type Order struct {
}

type OrderRepository interface {
	getById(int) Order
}

type orderRepository struct {
}

func (o *orderRepository) getById(id int) Order {
	return Order{}
}