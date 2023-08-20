package models

type Product struct {
	ID          int64
	ProductCode string
	Price       float64
	Stock       int
}

type CreateProduct struct {
	ProductCode string
	Price       float64
	Stock       int
}

type ProductPrice struct {
	Price float64
}

type ProductStock struct {
	Stock int
}

type UpdateProductStock struct {
	ProductCode string
	Stock       int
}
