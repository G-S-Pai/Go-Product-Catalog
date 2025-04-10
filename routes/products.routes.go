package routes

import (
	"github.com/g-s-pai/go-product-catalog/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func ProductRoutes(rg router.Party) {
	router := rg.Party("/products")

	router.Use(iris.Compression)
	router.Get("/", controllers.GetProducts)
	router.Get("/{id}", controllers.GetProductByID)
	router.Post("/", controllers.CreateProduct)
	router.Put("/{id}", controllers.UpdateProduct)
	router.Patch("/{id}", controllers.UpdateProduct)
	router.Delete("/{id}", controllers.DeleteProduct)
}
