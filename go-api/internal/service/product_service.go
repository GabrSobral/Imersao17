package service

import (
	"github.com/gabrsobral/imersao17/goapi/internal/database"
	"github.com/gabrsobral/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDb database.ProductDb
}

func NewProductService(productDb *database.ProductDb) *ProductService {
	return &ProductService{ProductDb: *productDb}
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := cs.ProductDb.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cs *ProductService) GetProductById(productId string) (*entity.Product, error) {
	product, err := cs.ProductDb.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	return product, err
}

func (cs *ProductService) GetProductByCategoryId(categoryId string) ([]*entity.Product, error) {
	products, err := cs.ProductDb.GetProductByCategoryId(categoryId)

	if err != nil {
		return nil, err
	}

	return products, err
}

func (cs *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := cs.ProductDb.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
