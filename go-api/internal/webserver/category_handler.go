package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/gabrsobral/imersao17/goapi/internal/entity"
	"github.com/gabrsobral/imersao17/goapi/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

func (wch *WebCategoryHandler) GetCategories(writter http.ResponseWriter, request *http.Request) {
	categories, error := wch.CategoryService.GetCategories()

	if error != nil {
		http.Error(writter, error.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(writter).Encode(categories)
}

func (wch *WebCategoryHandler) GetCategoryById(writter http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(writter, "Id is required", http.StatusBadRequest)
		return
	}

	category, error := wch.CategoryService.GetCategoryById(id)

	if error != nil {
		http.Error(writter, error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writter).Encode(category)
}

func (wch *WebCategoryHandler) CreateCategory(writter http.ResponseWriter, request *http.Request) {
	var category entity.Category

	err := json.NewDecoder(request.Body).Decode(&category)

	if err != nil {
		http.Error(writter, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.CategoryService.CreateCategory(category.Name)

	if err != nil {
		http.Error(writter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writter).Encode(result)
}
