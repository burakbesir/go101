package service

import (
	"fmt"
	"github.com/rahmanbesir/go101/10-example/model"
	"github.com/rahmanbesir/go101/10-example/repository"
)

type OrderService interface {
	UpdateOrderAsDelivered(int64) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (s *orderService) UpdateOrderAsDelivered(id int64) error {
	err, order := s.orderRepository.GetById(id)
	if err != nil {
		fmt.Printf("An error occurred when get order by id : %d", id)
		return err
	}

	setAsDelivered(order)

	return nil
}

func setAsDelivered(order model.Order) { // private function
	// update order
}
