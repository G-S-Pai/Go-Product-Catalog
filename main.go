package main

import (
	"github.com/g-s-pai/go-product-catalog/initializers"
	"github.com/g-s-pai/go-product-catalog/routes"

	"github.com/kataras/iris/v12"
)

var (
	app *iris.Application
	config initializers.Config
)

func init() {
	initializers.ConnectDB()

	app = iris.Default()
}

func main() {
	router := app.Party("/api/v1")

	routes.ProductRoutes(router)

	app.Listen(":8080")
}