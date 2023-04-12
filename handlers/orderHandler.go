package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharansharma94/product-order-system/order"
)

func GetOrderDetailsHandler(orderService *order.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID, _ := strconv.Atoi(c.Param("orderID"))

		order := orderService.GetOrderDetails(orderID)
		if order != nil {
			c.JSON(http.StatusOK, order)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Order not found."})
		}
	}
}

func UpdateOrderStatusHandler(orderService *order.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID, _ := strconv.Atoi(c.Param("orderID"))
		newStatus := order.Status(c.PostForm("status"))

		orderService.UpdateOrderStatus(orderID, newStatus)
		c.JSON(http.StatusOK, gin.H{"message": "Order status updated."})
	}
}

func UpdateDispatchDateHandler(orderService *order.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID, _ := strconv.Atoi(c.Param("orderID"))
		dispatchDate := c.PostForm("dispatch_date")

		orderService.UpdateDispatchDate(orderID, dispatchDate)
		c.JSON(http.StatusOK, gin.H{"message": "Dispatch date updated."})
	}
}
