package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharansharma94/product-order-system/order"
	"github.com/sharansharma94/product-order-system/product"
)

type PlaceOrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func GetCatalogueHandler(productService *product.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, productService.GetCatalogue())
	}
}

func PlaceOrderHandler(productService *product.ProductService, orderService *order.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData PlaceOrderRequest
		err := c.BindJSON(&requestData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
			return
		}

		orderID := productService.PlaceOrder(orderService, requestData.ProductID, requestData.Quantity)
		if orderID != -1 {
			c.JSON(http.StatusOK, gin.H{"order_id": orderID})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to place the order."})
		}
	}
}
