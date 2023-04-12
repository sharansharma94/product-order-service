package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharansharma94/product-order-system/order"
	"github.com/sharansharma94/product-order-system/product"
)

func GetCatalogueHandler(productService *product.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, productService.GetCatalogue())
	}
}

func PlaceOrderHandler(productService *product.ProductService, orderService *order.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID, _ := strconv.Atoi(c.PostForm("product_id"))
		quantity, _ := strconv.Atoi(c.PostForm("quantity"))

		orderID := productService.PlaceOrder(orderService, productID, quantity)
		if orderID != -1 {
			c.JSON(http.StatusOK, gin.H{"order_id": orderID})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to place the order."})
		}
	}
}
