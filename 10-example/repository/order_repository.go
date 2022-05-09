package repository

import (
	"errors"
	"github.com/rahmanbesir/go101/10-example/model"
)

// repository.NewOrderRepository

type OrderRepository interface {
	GetById(int64) (error, model.Order)
}

type orderRepository struct {
}

func NewOrderRepository() OrderRepository { // burada db için kullanacağımız paketi alıp initial edebiliriz
	return &orderRepository{}
}

func (r *orderRepository) GetById(id int64) (error, model.Order) {
	if id < 1 {
		return errors.New("An Error"), model.Order{}
	}
	return nil, model.Order{}
}
