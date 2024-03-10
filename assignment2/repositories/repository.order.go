package repositories

import (
	"assignment2/entities"
	"assignment2/schemas"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllOrderWithItems() ([]schemas.OrderOutput, error)
	InsertOrder(order entities.Order, items []entities.Item) error
	GetLatestOrderID() (int, error)
	GetLatestItemID() (int, error)
	DeleteOrderWithItems(orderID int) error
	UpdateOrder(orderID int, orderInput schemas.OrderInputUpdate) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetLatestOrderID() (int, error) {
    var latestID int
	result := r.db.Model(&entities.Order{}).Select("COALESCE(MAX(order_id), 0)").Scan(&latestID)
    if result.Error != nil {
        return 0, result.Error
    }
    return latestID, nil
}

func (r *repository) GetLatestItemID() (int, error) {
    var latestID int
	result := r.db.Model(&entities.Item{}).Select("COALESCE(MAX(item_id), 0)").Scan(&latestID)
    if result.Error != nil {
        return 0, result.Error
    }
    return latestID, nil
}

func (r *repository) InsertOrder(order entities.Order, items []entities.Item) error {
	// Mulai transaksi
	tx := r.db.Begin()

	// Sisipkan data pesanan
	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Dapatkan order_id dari pesanan yang baru saja disisipkan
	orderID := order.OrderID

	// Sisipkan data item untuk pesanan yang baru saja disisipkan
	for i := range items {
		items[i].OrderID = orderID
		err = tx.Create(&items[i]).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaksi
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repository) UpdateOrder(orderID int, orderInput schemas.OrderInputUpdate) error {
	// Example using ORM
	var order entities.Order
	if err := r.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return errors.New("failed to find order")
	}

	// Update order details
	order.CustomerName = orderInput.CustomerName
	order.OrderedAt = orderInput.OrderedAt

	// Save the updated order
	if err := r.db.Where("order_id = ?", orderID).Save(&order).Error; err != nil {
		fmt.Println(err)
		return errors.New("failed to update order")
	}

	// Update item details
	for _, itemInput := range orderInput.Items {
		var item entities.Item
		if err := r.db.Where("item_id = ?", itemInput.LineItemID).Find(&item).Error; err != nil {
			return errors.New("failed to find item")
		}

		// Update item details
		item.ItemCode = itemInput.ItemCode
		item.Description = itemInput.Description
		item.Quantity = itemInput.Quantity
		
		// Save the updated item
		if err := r.db.Where("item_id = ?", itemInput.LineItemID).Omit("item_id").Updates(&item).Error; err != nil {
			return errors.New("failed to update item")
		}
	}

	return nil
}

func (r *repository) GetAllOrderWithItems() ([]schemas.OrderOutput, error) {
    var orders []entities.Order
    err := r.db.Find(&orders).Error
    if err != nil {
        return nil, err
    }

    var orderOutputs []schemas.OrderOutput
    for _, order := range orders {
        var items []entities.Item
        err := r.db.Where("order_id = ?", order.OrderID).Find(&items).Error
        if err != nil {
            return nil, err
        }
        orderOutput := schemas.OrderOutput{
            OrderID:      order.OrderID,
            OrderedAt:    order.OrderedAt,
            CustomerName: order.CustomerName,
            Items:        items,
        }
        orderOutputs = append(orderOutputs, orderOutput)
    }

    return orderOutputs, nil
}

func (r *repository) DeleteOrderWithItems(orderID int) error {
    // Mulai transaksi
    tx := r.db.Begin()

    // Hapus semua item terkait dengan pesanan
    if err := tx.Where("order_id = ?", orderID).Delete(&entities.Item{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Hapus pesanan
    if err := tx.Where("order_id = ?", orderID).Delete(&entities.Order{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Commit transaksi
    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return err
    }

    return nil
}


