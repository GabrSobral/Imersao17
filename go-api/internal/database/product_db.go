package database

import (
	"database/sql"

	"github.com/gabrsobral/imersao17/goapi/internal/entity"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (pd *ProductDb) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, category_id, image_url FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.ImageUrl); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (cd *ProductDb) GetProductById(productId string) (*entity.Product, error) {
	var product entity.Product

	err := cd.db.
		QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", productId).
		Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.ImageUrl)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDb) GetProductByCategoryId(categoryId string) ([]*entity.Product, error) {
	rows, err := pd.db.
		Query("SELECT id, name, description, price, category_id, image_url FROM products WHERE category_id = ?", categoryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.ImageUrl); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDb) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)",
		product.Id, product.Name, product.Description, product.Price, product.CategoryId, product.ImageUrl)

	if err != nil {
		return nil, err
	}

	return product, nil
}
