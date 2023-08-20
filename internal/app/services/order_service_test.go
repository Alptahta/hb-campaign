package services

import (
	"errors"
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/faketime"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

var (
	mockCreateOrderRequest = models.CreateOrderRequest{
		ProductCode: "P1",
		Quantity:    100,
	}
	fakeProductPrice = models.ProductPrice{
		Price: 100.0,
	}

	fakeProductStock = models.ProductStock{
		Stock: 100,
	}
)

func Test_CreateOrder(t *testing.T) {
	t.Run("It should return an error when GetProductPriceByProductCode function from product service layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockOrderRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		mockOrderService := NewMockProductServiceI(mockCtrl)
		ps := NewOrderService(mockFakeTime, mockOrderService, mockRepository)

		mockOrderService.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(nil, errors.New("fake error")).
			Times(1)

		err := ps.CreateOrder(mockCreateOrderRequest)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should return an error when GetProductStockByProductName function from product service layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockOrderRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		mockOrderService := NewMockProductServiceI(mockCtrl)
		ps := NewOrderService(mockFakeTime, mockOrderService, mockRepository)

		mockOrderService.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(&fakeProductPrice, nil).
			Times(1)

		mockOrderService.EXPECT().
			GetProductStockByProductName(gomock.Any()).
			Return(nil, errors.New("fake error")).
			Times(1)

		err := ps.CreateOrder(mockCreateOrderRequest)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should return an error when UpdateProductStockByProductName function from product service layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockOrderRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		mockOrderService := NewMockProductServiceI(mockCtrl)
		ps := NewOrderService(mockFakeTime, mockOrderService, mockRepository)

		mockOrderService.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(&fakeProductPrice, nil).
			Times(1)

		mockOrderService.EXPECT().
			GetProductStockByProductName(gomock.Any()).
			Return(&fakeProductStock, nil).
			Times(1)

		mockOrderService.EXPECT().
			UpdateProductStockByProductName(gomock.Any()).
			Return(errors.New("fake error")).
			Times(1)

		err := ps.CreateOrder(mockCreateOrderRequest)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should return an error when CreateOrder function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockOrderRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		mockOrderService := NewMockProductServiceI(mockCtrl)
		ps := NewOrderService(mockFakeTime, mockOrderService, mockRepository)

		mockOrderService.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(&fakeProductPrice, nil).
			Times(1)

		mockOrderService.EXPECT().
			GetProductStockByProductName(gomock.Any()).
			Return(&fakeProductStock, nil).
			Times(1)

		mockOrderService.EXPECT().
			UpdateProductStockByProductName(gomock.Any()).
			Return(nil).
			Times(1)

		mockRepository.EXPECT().
			CreateOrder(gomock.Any()).
			Return(createdCampaignID, errors.New("fake error")).
			Times(1)

		err := ps.CreateOrder(mockCreateOrderRequest)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should not return any error when CreateOrder function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockOrderRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		mockOrderService := NewMockProductServiceI(mockCtrl)
		ps := NewOrderService(mockFakeTime, mockOrderService, mockRepository)

		mockOrderService.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(&fakeProductPrice, nil).
			Times(1)

		mockOrderService.EXPECT().
			GetProductStockByProductName(gomock.Any()).
			Return(&fakeProductStock, nil).
			Times(1)

		mockOrderService.EXPECT().
			UpdateProductStockByProductName(gomock.Any()).
			Return(nil).
			Times(1)

		mockRepository.EXPECT().
			CreateOrder(gomock.Any()).
			Return(createdCampaignID, nil).
			Times(1)

		err := ps.CreateOrder(mockCreateOrderRequest)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
	})
}
