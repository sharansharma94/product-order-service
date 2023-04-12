package order

type Status string

const (
	StatusPlaced     Status = "Placed"
	StatusDispatched Status = "Dispatched"
	StatusCompleted  Status = "Completed"
	StatusCancelled  Status = "Cancelled"
  DiscountRate = 0.9
)

type Order struct{
  ID int
  OrderValue float64
  DispatchDate string
  OrderStatus Status
  ProdQuantity int
  PremiumProduct int
}

type OrderService struct {
  orders []Order
  orderID int
}

func NewOrderService() *OrderService {
  return &OrderService{
    orders: []Order{},
    orderID: 1,
  }
}

func (os *OrderService) GetOrderDetails(orderID int) *Order {
  for i :=  range os.orders{
    if os.orders[i].ID == orderID{
      return &os.orders[i]
    }
  }
  return nil
}

func (os *OrderService) ApplyDiscount(orderID int){
  order := os.GetOrderDetails(orderID)
  if order != nil && order.PremiumProduct >= 3 {
    order.OrderValue *= DiscountRate
  }
}

func (os *OrderService) UpdateOrderStatus(orderID int, newStatus Status){

  order := os.GetOrderDetails(orderID)
  if order != nil {
    order.OrderStatus = newStatus
  }
}

func (os *OrderService) UpdateDispatchDate(orderID int, dispatchDate string){

  order := os.GetOrderDetails(orderID)
  if order != nil && order.OrderStatus == StatusDispatched{
    order.DispatchDate = dispatchDate
  }
}

func (os *OrderService) AddOrder(productID int, orderValue float64, prodQuantity int, premiumProduct int) int {
	newOrder := Order{
		ID:             os.orderID,
		OrderValue:     orderValue,
		DispatchDate:   "",
		OrderStatus:    StatusPlaced,
		ProdQuantity:   prodQuantity,
		PremiumProduct: premiumProduct,
	}
	os.orders = append(os.orders, newOrder)
	os.orderID++
	return newOrder.ID
}


