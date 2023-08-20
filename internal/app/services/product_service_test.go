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
	mockCreateProductData = models.CreateProduct{
		ProductCode: "P1",
		Price:       100.0,
		Stock:       100,
	}

	mockProductCode = "P1"

	lastCreatedID int64 = 1

	fakeProductResponse = models.Product{
		ID:          1,
		ProductCode: "P1",
		Price:       100.0,
		Stock:       100,
	}

	fakeProductPriceResponse = models.ProductPrice{
		Price: 100.0,
	}

	fakeProductStockResponse = models.ProductStock{
		Stock: 100,
	}

	fakeUpdateProductStock = models.UpdateProductStock{
		ProductCode: "P1",
		Stock:       100,
	}
)

func Test_CreateProduct(t *testing.T) {
	t.Run("It should return an error when CreateProduct function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			CreateProduct(gomock.Any()).
			Return(lastCreatedID, errors.New("fake error")).
			Times(1)

		err := ps.CreateProduct(mockCreateProductData)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should not return any error when CreateProduct function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			CreateProduct(gomock.Any()).
			Return(lastCreatedID, nil).
			Times(1)

		err := ps.CreateProduct(mockCreateProductData)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
	})
}

func Test_GetProductByProductCode(t *testing.T) {
	t.Run("It should return an error when GetProductByProductCode function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetProductByProductCode(gomock.Any()).
			Return(models.Product{}, errors.New("fake error")).
			Times(1)

		product, err := ps.GetProductByProductCode(mockProductCode)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
		if product != nil {
			t.Fatalf("expected nil but got %v", product)
		}
	})

	t.Run("It should not return any error when GetProductByProductCode function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetProductByProductCode(gomock.Any()).
			Return(fakeProductResponse, nil).
			Times(1)

		product, err := ps.GetProductByProductCode(mockProductCode)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
		if product == nil {
			t.Fatal("expected non nil but got nil")
		}
	})
}

func Test_GetProductPriceByProductCode(t *testing.T) {
	t.Run("It should return an error when GetProductPriceByProductCode function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(models.ProductPrice{}, errors.New("fake error")).
			Times(1)

		product, err := ps.GetProductPriceByProductCode(mockProductCode)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
		if product != nil {
			t.Fatalf("expected nil but got %v", product)
		}
	})

	t.Run("It should not return any error when GetProductPriceByProductCode function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetProductPriceByProductCode(gomock.Any()).
			Return(fakeProductPriceResponse, nil).
			Times(1)

		product, err := ps.GetProductPriceByProductCode(mockProductCode)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
		if product == nil {
			t.Fatal("expected non nil but got nil")
		}
	})
}

func Test_GetProductStockByProductName(t *testing.T) {
	t.Run("It should return an error when GetProductStockByProductName function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetProductStockByProductName(gomock.Any()).
			Return(models.ProductStock{}, errors.New("fake error")).
			Times(1)

		product, err := ps.GetProductStockByProductName(mockProductCode)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
		if product != nil {
			t.Fatalf("expected nil but got %v", product)
		}
	})

	t.Run("It should not return any error when GetProductStockByProductName function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			GetProductStockByProductName(gomock.Any()).
			Return(fakeProductStockResponse, nil).
			Times(1)

		product, err := ps.GetProductStockByProductName(mockProductCode)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
		if product == nil {
			t.Fatal("expected non nil but got nil")
		}
	})
}

func Test_UpdateProductStockByProductName(t *testing.T) {
	t.Run("It should return an error when UpdateProductStockByProductName function from repository layer returns an error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			UpdateProductStockByProductName(gomock.Any()).
			Return(errors.New("fake error")).
			Times(1)

		err := ps.UpdateProductStockByProductName(fakeUpdateProductStock)
		if err == nil {
			t.Fatalf("expected error but had no error")
		}
	})

	t.Run("It should not return any error when UpdateProductStockByProductName function from repository layer does not returns any error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()
		mockRepository := repositories.NewMockProductRepositoryI(mockCtrl)
		mockFakeTime := faketime.NewMockTimeInterface(mockCtrl)
		ps := NewProductService(mockFakeTime, mockRepository)

		mockRepository.EXPECT().
			UpdateProductStockByProductName(gomock.Any()).
			Return(nil).
			Times(1)

		err := ps.UpdateProductStockByProductName(fakeUpdateProductStock)
		if err != nil {
			t.Fatalf("did not expect any error but got %v", err)
		}
	})
}
