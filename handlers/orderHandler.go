package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sharansharma94/product-order-system/order"
)

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

type UpdateOrderDispatchDateRequest struct {
	DispatchDate string `json:"dispatch_date"`
}

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
		var requestData UpdateOrderStatusRequest
		err := c.BindJSON(&requestData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
			return
		}

		newStatus := order.Status(requestData.Status)

		orderService.UpdateOrderStatus(orderID, newStatus)
		c.JSON(http.StatusOK, gin.H{"message": "Order status updated."})
	}
}
func UpdateDispatchDateHandler(orderService *order.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID, _ := strconv.Atoi(c.Param("orderID"))
		var requestData UpdateOrderDispatchDateRequest
		err := c.BindJSON(&requestData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
			return
		}

		_, err = time.Parse("2006-01-02", requestData.DispatchDate)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid date format"})
			return
		}

		orderService.UpdateDispatchDate(orderID, requestData.DispatchDate)
		c.JSON(http.StatusOK, gin.H{"message": "Dispatch date updated."})
	}
}

// func UpdateDispatchDateHandler(orderService *order.OrderService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		orderID, _ := strconv.Atoi(c.Param("orderID"))
// 		dispatchDate := c.PostForm("dispatch_date")

// 		orderService.UpdateDispatchDate(orderID, dispatchDate)
// 		c.JSON(http.StatusOK, gin.H{"message": "Dispatch date updated."})
// 	}
// }
