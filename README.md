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
