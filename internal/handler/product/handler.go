package product

import (
	"net/http"

	"github.com/eulabs/back-end/api-product/internal/config/logger"
	"github.com/eulabs/back-end/api-product/pkg/model"
	"github.com/eulabs/back-end/api-product/pkg/service/product"
	"github.com/labstack/echo/v4"
)

func getAllProducts(service product.ProductServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		list_products := service.GetAll(c.Request().Context())
		return c.JSON(http.StatusOK, list_products)
	}
}

func getProduct(service product.ProductServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		externalID := c.Request().Header.Get("id")

		if externalID == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
		}

		product := service.GetByExternalID(c.Request().Context(), externalID)
		if product.ID == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}

		return c.JSON(http.StatusOK, product)
	}
}

func createProduct(service product.ProductServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := model.Product{}

		if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
		}

		prd, err := model.NewProduct(&product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product data"})
		}

		logger.Info(prd.Name)
		if prd.Name == " " || prd.Name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product name is required"})
		}

		if prd.Quantity <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product quantity is required"})
		}

		if prd.Price <= 0.0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product price is required"})
		}

		product = *prd

		external_id := service.Create(c.Request().Context(), &product)
		if external_id == "" {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert product"})
		}

		product = *service.GetByExternalID(c.Request().Context(), external_id)

		return c.JSON(http.StatusOK, product)
	}
}

func updateProduct(service product.ProductServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {

		request_to_update_product := model.Product{}

		if err := c.Bind(&request_to_update_product); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
		}

		if request_to_update_product.ExternalID == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
		}
		externalID := request_to_update_product.ExternalID

		old_product := service.GetByExternalID(c.Request().Context(), externalID)
		if old_product.ID == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}

		rows_affected := service.Update(c.Request().Context(), old_product.ExternalID, &request_to_update_product)
		if rows_affected == 0 {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
		}

		request_to_update_product.ExternalID = old_product.ExternalID

		return c.JSON(http.StatusOK, request_to_update_product)
	}
}

func deleteProduct(service product.ProductServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		externalID := c.Request().Header.Get("id")

		if externalID == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
		}

		old_product := service.GetByExternalID(c.Request().Context(), externalID)
		if old_product.ExternalID == "" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}

		rows_affected := service.Delete(c.Request().Context(), old_product.ExternalID)
		if rows_affected == 0 {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
	}
}
