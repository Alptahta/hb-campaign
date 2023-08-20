package services

import (
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/faketime"
)

type OrderServiceI interface {
	CreateOrder(co models.CreateOrderRequest) error
}

type OrderService struct {
	t  faketime.TimeInterface
	ps ProductServiceI
	or repositories.OrderRepositoryI
}

func NewOrderService(t faketime.TimeInterface, ps ProductServiceI, or repositories.OrderRepositoryI) *OrderService {
	return &OrderService{
		t:  t,
		ps: ps,
		or: or,
	}
}

func (os OrderService) CreateOrder(co models.CreateOrderRequest) error {
	price, err := os.ps.GetProductPriceByProductCode(co.ProductCode)
	if err != nil {
		return err
	}

	var totalPrice float64 = price.Price * float64(co.Quantity)
	orderDTO := models.CreateOrderDTO{
		ProductCode: co.ProductCode,
		Quantity:    co.Quantity,
		Price:       totalPrice,
	}

	stock, err := os.ps.GetProductStockByProductName(co.ProductCode)
	if err != nil {
		return err
	}
	newStock := stock.Stock - co.Quantity

	updateProductStock := models.UpdateProductStock{
		ProductCode: co.ProductCode,
		Stock:       newStock,
	}

	err = os.ps.UpdateProductStockByProductName(updateProductStock)
	if err != nil {
		return err
	}

	_, err = os.or.CreateOrder(orderDTO)
	if err != nil {
		return err
	}
	return nil
}
