package main

import (
	"fmt"

	"github.com/sharansharma94/product-order-system/order"
	"github.com/sharansharma94/product-order-system/product"
)

func main() {
	productService := product.NewProductService()
	orderService := order.NewOrderService()

	catalogue := productService.GetCatalogue()
  fmt.Println("Product Catalogue:")

  for _, product := range catalogue{
    fmt.Printf("ID: %d, Name: %s, Category: %s, Price %.2f, Availability: %d\n", product.ID, product.Name, product.Category, product.Price, product.Availability) 
  }

  // Simulate placing an order
	orderID := productService.PlaceOrder(orderService,1, 5) // Product with ID 1 and quantity 5
	fmt.Printf("Order placed with ID: %d\n", orderID)

	// Simulate applying a discount on an order
	orderService.ApplyDiscount(orderID)

	// Simulate updating the order status
	orderService.UpdateOrderStatus(orderID, order.StatusDispatched)

	// Simulate updating the dispatch date
	orderService.UpdateDispatchDate(orderID, "2023-04-15")

	// Get order details
	orderDetails := orderService.GetOrderDetails(orderID)
	fmt.Printf("Order ID: %d, Order Value: %.2f, Dispatch Date: %s, Order Status: %s, Product Quantity: %d\n", orderDetails.ID, orderDetails.OrderValue, orderDetails.DispatchDate, orderDetails.OrderStatus, orderDetails.ProdQuantity)
}
