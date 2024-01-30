package entity

import "github.com/google/uuid"

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		Id:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  string  `json:"category_id"`
	ImageUrl    string  `json:"image_url"`
}

func NewProduct(name, description, categoryId, imageUrl string, price float64) *Product {
	return &Product{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryId:  categoryId,
		ImageUrl:    imageUrl,
	}
}
