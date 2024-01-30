package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/gabrsobral/imersao17/goapi/internal/entity"
	"github.com/gabrsobral/imersao17/goapi/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wch *WebProductHandler) GetProducts(writter http.ResponseWriter, request *http.Request) {
	products, error := wch.ProductService.GetProducts()

	if error != nil {
		http.Error(writter, error.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(writter).Encode(products)
}

func (wch *WebProductHandler) GetProductById(writter http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(writter, "Id is required", http.StatusBadRequest)
		return
	}

	category, error := wch.ProductService.GetProductById(id)

	if error != nil {
		http.Error(writter, error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writter).Encode(category)
}

func (wch *WebProductHandler) GetProductByCategoryId(writter http.ResponseWriter, request *http.Request) {
	categoryId := chi.URLParam(request, "categoryId")

	if categoryId == "" {
		http.Error(writter, "Category Id is required", http.StatusBadRequest)
		return
	}

	category, error := wch.ProductService.GetProductByCategoryId(categoryId)

	if error != nil {
		http.Error(writter, error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writter).Encode(category)
}

func (wch *WebProductHandler) CreateProduct(writter http.ResponseWriter, request *http.Request) {
	var product entity.Product

	err := json.NewDecoder(request.Body).Decode(&product)

	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.ProductService.CreateProduct(
		product.Name,
		product.Description,
		product.CategoryId,
		product.ImageUrl,
		product.Price)

	if err != nil {
		http.Error(writter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writter).Encode(result)
}
