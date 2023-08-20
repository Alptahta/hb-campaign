package services

import (
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/faketime"
)

type ProductServiceI interface {
	CreateProduct(cp models.CreateProduct) error
	GetProductByProductCode(string) (*models.Product, error)
	GetProductPriceByProductCode(productCode string) (*models.ProductPrice, error)
	GetProductStockByProductName(productCode string) (*models.ProductStock, error)
	UpdateProductStockByProductName(models.UpdateProductStock) error
	Update() error
}

type ProductService struct {
	t  faketime.TimeInterface
	pr repositories.ProductRepositoryI
	cs CampaignServiceI
}

func NewProductService(t faketime.TimeInterface, pr repositories.ProductRepositoryI, cs CampaignServiceI) *ProductService {
	return &ProductService{
		t:  t,
		pr: pr,
		cs: cs,
	}
}

func (p ProductService) CreateProduct(cp models.CreateProduct) error {
	_, err := p.pr.CreateProduct(cp)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductService) GetProductByProductCode(productCode string) (*models.Product, error) {
	product, err := p.pr.GetProductByProductCode(productCode)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p ProductService) GetProductPriceByProductCode(productCode string) (*models.ProductPrice, error) {
	productPrice, err := p.pr.GetProductPriceByProductCode(productCode)
	if err != nil {
		return nil, err
	}
	return &productPrice, nil
}

func (p ProductService) GetProductStockByProductName(productCode string) (*models.ProductStock, error) {
	productStock, err := p.pr.GetProductStockByProductName(productCode)
	if err != nil {
		return nil, err
	}
	return &productStock, nil
}

func (p ProductService) UpdateProductStockByProductName(up models.UpdateProductStock) error {
	err := p.pr.UpdateProductStockByProductName(up)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductService) Update() error {
	campaigns, err := p.cs.GetAllActiveCampaigns()
	if err != nil {
		return err
	}
	for _, campaign := range campaigns {
		price, err := p.pr.GetProductPriceByProductCode(campaign.ProductCode)
		if err != nil {
			return err
		}
		err = p.pr.UpdateProductPriceByProductCode(campaign.ProductCode, price.Price-float64(10))
		if err != nil {
			return err
		}
	}

	campaigns, err = p.cs.GetAllEndedCampaigns()
	if err != nil {
		return err
	}
	for _, campaign := range campaigns {
		price, err := p.pr.GetProductPriceByProductCode(campaign.ProductCode)
		if err != nil {
			return err
		}
		err = p.pr.UpdateProductPriceByProductCode(campaign.ProductCode, price.Price)
		if err != nil {
			return err
		}
	}

	return nil
}
