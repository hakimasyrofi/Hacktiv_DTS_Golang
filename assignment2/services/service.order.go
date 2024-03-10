package services

import (
	"assignment2/entities"
	"assignment2/repositories"
	"assignment2/schemas"
)

type Service interface {
	CreateOrder(order entities.Order, items []entities.Item) (error)
	GetAllOrders() ([]schemas.OrderOutput, error)
	DeleteOrderWithItems(orderID int) error
	UpdateOrder(orderID int, orderInput schemas.OrderInputUpdate) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllOrders() ([]schemas.OrderOutput, error) {
	items, err := s.repository.GetAllOrderWithItems()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *service) CreateOrder(order entities.Order, items []entities.Item) error {
	latestOrderID, err := s.repository.GetLatestOrderID()
	if err != nil {
		return err
	}
	latestItemID, err := s.repository.GetLatestItemID()
	if err != nil {
		return err
	}
	latestOrderID++
	order.OrderID = latestOrderID
	
	for i := range items {
		latestItemID++
		items[i].ItemID = latestItemID
	}
	err = s.repository.InsertOrder(order, items)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteOrderWithItems(orderID int) error {
	err := s.repository.DeleteOrderWithItems(orderID)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateOrder(orderID int, orderInput schemas.OrderInputUpdate) error {
	err := s.repository.UpdateOrder(orderID, orderInput)

	if err != nil {
		return err
	}

	return nil
}