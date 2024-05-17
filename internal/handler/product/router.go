package product

import (
	"github.com/eulabs/back-end/api-product/pkg/service/product"
	"github.com/labstack/echo/v4"
)

func RegisterProductAPIHandlers(e *echo.Echo, service product.ProductServiceInterface) {
	v1 := e.Group("/api/v1/prd")

	v1.POST("/add", createProduct(service))
	v1.PATCH("/update", updateProduct(service))
	v1.DELETE("/delete", deleteProduct(service))
	v1.GET("/getbyid", getProduct(service))
	v1.GET("/all", getAllProducts(service))
}
