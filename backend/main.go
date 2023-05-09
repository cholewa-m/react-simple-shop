package main

import (
	"net/http"

	"CRUDgo/controller"
	"CRUDgo/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Dodaj middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Product{})

	productController := controller.NewProductController(db)
	paymentController := controller.NewPaymentController()

	e.GET("/", hello)
	e.GET("/products", productController.GetAllProducts)
	e.GET("/products/:id", productController.GetProductById)
	e.POST("/products", productController.CreateProduct)

	e.POST("/payments", paymentController.CreatePayment) // Dodajemy trasę POST dla płatności

	e.Logger.Fatal(e.Start(":1323"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK,
		"Hello! "+
			"\n Get all products: /products "+
			"\n Get product by id: /products/:id "+
			"\n Post product with name and price: /products?name=[name]&price=[price]")
}
