package repositories

import (
	"database/sql"
	"fmt"
	"hb-campaign/internal/app/models"
)

type ProductRepositoryI interface {
	CreateProduct(cp models.CreateProduct) (int64, error)
	GetProductByProductCode(productCode string) (models.Product, error)
	GetProductPriceByProductCode(productCode string) (models.ProductPrice, error)
	GetProductStockByProductName(productCode string) (models.ProductStock, error)
	UpdateProductStockByProductName(up models.UpdateProductStock) error
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) CreateProduct(p models.CreateProduct) (int64, error) {
	result, err := pr.db.Exec("INSERT INTO products (product_code, price, stock) VALUES (?, ?, ?)", p.ProductCode, p.Price, p.Stock)
	if err != nil {
		return 0, fmt.Errorf("addProduct: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addProduct: %v", err)
	}
	return id, nil
}

func (pr *ProductRepository) GetProductByProductCode(productCode string) (models.Product, error) {
	var product models.Product
	row := pr.db.QueryRow("SELECT * FROM products WHERE product_code = ?", productCode)
	if err := row.Scan(&product.ID, &product.ProductCode, &product.Price, &product.Stock); err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("productByProductCode %s: no such product", productCode)
		}
		return product, fmt.Errorf("productByProductCode %s: %v", productCode, err)
	}
	return product, nil
}

func (pr *ProductRepository) GetProductPriceByProductCode(productCode string) (models.ProductPrice, error) {
	var productPrice models.ProductPrice
	row := pr.db.QueryRow("SELECT price FROM products WHERE product_code = ?", productCode)
	if err := row.Scan(&productPrice.Price); err != nil {
		if err == sql.ErrNoRows {
			return productPrice, fmt.Errorf("productByProductCode %s: no such product", productCode)
		}
		return productPrice, fmt.Errorf("productByProductCode %s: %v", productCode, err)
	}
	return productPrice, nil
}

func (pr *ProductRepository) GetProductStockByProductName(productCode string) (models.ProductStock, error) {
	var productStock models.ProductStock
	row := pr.db.QueryRow("SELECT stock FROM products WHERE product_code = ?", productCode)
	if err := row.Scan(&productStock.Stock); err != nil {
		if err == sql.ErrNoRows {
			return productStock, fmt.Errorf("productByProductCode %s: no such product", productCode)
		}
		return productStock, fmt.Errorf("productByProductCode %s: %v", productCode, err)
	}
	return productStock, nil
}

func (pr *ProductRepository) UpdateProductStockByProductName(up models.UpdateProductStock) error {
	res, err := pr.db.Exec("UPDATE products SET stock=? where product_code=?", up.Stock, up.ProductCode)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
