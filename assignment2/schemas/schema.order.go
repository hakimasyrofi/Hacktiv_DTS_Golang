package schemas

import (
	"assignment2/entities"
	"time"
)

type OrderInput struct {
	OrderedAt    time.Time     
	CustomerName string          
	Items        []entities.Item
}

type OrderInputUpdate struct {
	OrderedAt    time.Time     
	CustomerName string          
	Items        []entities.UpdatingItem
}

type OrderOutput struct {
	OrderID      int
	OrderedAt    time.Time     
	CustomerName string          
	Items        []entities.Item
}