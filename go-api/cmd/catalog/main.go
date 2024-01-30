package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gabrsobral/imersao17/goapi/internal/database"
	"github.com/gabrsobral/imersao17/goapi/internal/service"
	"github.com/gabrsobral/imersao17/goapi/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDb := database.NewCategoryDb(db)
	categoryService := service.NewCategoryService(categoryDb)

	productDb := database.NewProductDb(db)
	productService := service.NewProductService(productDb)

	webCategoryHandler := webserver.NewCategoryHandler(categoryService)
	webProductHandler := *webserver.NewProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)
	c.Get("/category/{id}", webCategoryHandler.GetCategoryById)

	c.Get("/product", webProductHandler.GetProducts)
	c.Post("/product", webProductHandler.CreateProduct)
	c.Get("/product/{id}", webProductHandler.GetProductById)
	c.Get("/product/category/{categoryId}", webProductHandler.GetProductByCategoryId)

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", c)
}
