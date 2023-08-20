package repositories

import (
	"database/sql"
	"fmt"
	"hb-campaign/internal/app/models"
)

type OrderRepositoryI interface {
	CreateOrder(cc models.CreateOrderDTO) (int64, error)
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (pr *OrderRepository) CreateOrder(o models.CreateOrderDTO) (int64, error) {
	result, err := pr.db.Exec("INSERT INTO orders (product_code, quantity, price) VALUES (?, ?, ?)",
		o.ProductCode, o.Quantity, o.Price)
	if err != nil {
		return 0, fmt.Errorf("addOrder: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addOrder: %v", err)
	}
	return id, nil
}
