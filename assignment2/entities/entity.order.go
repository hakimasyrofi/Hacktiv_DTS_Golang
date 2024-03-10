package entities

import "time"

type Order struct {
	OrderID      int
	CustomerName string
	OrderedAt    time.Time
}
