package controllers

import (
	"github.com/g-s-pai/go-product-catalog/initializers"
	"github.com/g-s-pai/go-product-catalog/models"

	"github.com/kataras/iris/v12"
)

func GetProducts(ctx iris.Context) {
	var products []models.Product
	if err := initializers.DB.Find(&products).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(products)
}

func GetProductByID(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var product models.Product
	if err := initializers.DB.First(&product, "id = ?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Product not found"})
		return
	}
	ctx.JSON(product)
}

func CreateProduct(ctx iris.Context) {
	var product models.Product
	if err := ctx.ReadJSON(&product); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&product).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(product)
}

func UpdateProduct(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var product models.Product
	var checkProduct models.Product

	if err := initializers.DB.First(&checkProduct, "id = ?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Product not found"})
		return
	}

    if err := ctx.ReadJSON(&product); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

    if err := initializers.DB.Save(&product).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(product)
}

func DeleteProduct(ctx iris.Context) {
	id := ctx.Params().Get("id")
	if err := initializers.DB.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
}
