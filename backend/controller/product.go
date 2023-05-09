package controller

import (
	"net/http"
	"strconv"

	"CRUDgo/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db}
}

func (controller *ProductController) GetAllProducts(ctx echo.Context) error {
	var products []model.Product
	result := controller.db.Find(&products)
	if result.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, result.Error)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (controller *ProductController) GetProductById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid product ID")
	}
	var product model.Product
	result := controller.db.First(&product, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, "Product not found")
	}
	return ctx.JSON(http.StatusOK, product)
}

func (controller *ProductController) CreateProduct(ctx echo.Context) error {

	name := ctx.QueryParam("name")
	price, err := strconv.ParseFloat(ctx.QueryParam("price"), 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid product price!")
	}

	controller.db.Create(&model.Product{Name: name, Price: price})

	return ctx.JSON(http.StatusCreated, "Product created")
}
