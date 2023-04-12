package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sharansharma94/product-order-system/handlers"
	"github.com/sharansharma94/product-order-system/order"
	"github.com/sharansharma94/product-order-system/product"
)

func main() {

	// Simulate()
	productService := product.NewProductService()
	orderService := order.NewOrderService()

	router := gin.Default()

	router.GET("/catalogue", handlers.GetCatalogueHandler(productService))
	router.POST("/order", handlers.PlaceOrderHandler(productService, orderService))
	router.GET("/order/:orderID", handlers.GetOrderDetailsHandler(orderService))
	router.PUT("/order/:orderID/status", handlers.UpdateOrderStatusHandler(orderService))
	router.PUT("/order/:orderID/dispatch", handlers.UpdateDispatchDateHandler(orderService))

	router.Run(":8080")
}
