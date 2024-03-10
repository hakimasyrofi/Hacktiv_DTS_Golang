package entities

type Item struct {
	ItemID      int
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
}

type UpdatingItem struct {
	Item
	LineItemID int
}
