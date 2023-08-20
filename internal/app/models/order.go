package models

type Order struct {
	ID          int64
	ProductCode string
	Quantity    int
	Price       float64
}

type CreateOrderRequest struct {
	ProductCode string
	Quantity    int
}

type CreateOrderDTO struct {
	ProductCode string
	Quantity    int
	Price       float64
}
