# product-order-service

Below are the two services and their corresponding operations based on the requirements provided:

### ProductService

- `getCatalogue()`: Retrieves a list of products with their availability, price, and category.
- `placeOrder(orderService, productID, quantity)`: Places an order for a particular product with the specified quantity (max 10).
- `updateProductCatalogue(productID, quantity)`: Updates the product catalogue after an order is placed by reducing the available quantity of the product.

### OrderService

- `getOrderDetails(orderID)`: Retrieves information about the order, such as orderValue, dispatchDate, orderStatus, and prodQuantity.
- `applyDiscount(orderID)`: Checks if the order contains 3 different premium products and applies a 10% discount to the order value if the condition is met.
- `updateOrderStatus(orderID, newStatus)`: Updates the orderStatus for a particular order.
- `updateDispatchDate(orderID, dispatchDate)`: Updates the dispatchDate for an order when the orderStatus is changed to 'Dispatched'.

**Product categories:** Premium, Regular, and Budget

**Order status values:** Placed, Dispatched, Completed, and Cancelled


## Commands to test

1. Get the product catalogue:

```sh
curl -X GET "http://localhost:8080/catalogue"
```
2. Place an order (example: product_id=1 and quantity=3):

```sh 
curl -X POST "http://localhost:8080/order" -d "product_id=1&quantity=3"
```

3. Get order details (example: orderID=1):
```sh
curl -X GET "http://localhost:8080/order/1"
```
4. Update order status (example: orderID=1 and new status=Dispatched):
```sh
curl -X PUT "http://localhost:8080/order/1/status" -d "status=Dispatched"
```
5. Update dispatch date (example: orderID=1 and dispatch_date=2023-04-20):
```sh
curl -X PUT "http://localhost:8080/order/1/dispatch" -d "dispatch_date=2023-04-20"
```
