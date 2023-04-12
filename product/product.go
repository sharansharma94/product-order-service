package product

import "github.com/sharansharma94/product-order-system/order"

type Product struct {
	ID           int
	Name         string
	Category     string
	Price        float64
	Availability int
}

type ProductService struct {
	catalogue []Product
}

func NewProductService() *ProductService {
	return &ProductService{
		catalogue: []Product{
			{ID: 1, Name: "Product 1", Category: "Premium", Price: 100.0, Availability: 20},
			{ID: 2, Name: "Product 2", Category: "Regular", Price: 50.0, Availability: 30},
			{ID: 3, Name: "Product 3", Category: "Budget", Price: 20.0, Availability: 50},
		},
	}
}

func (ps *ProductService) GetCatalogue() []Product {
	return ps.catalogue
}

func (ps *ProductService) PlaceOrder(orderService *order.OrderService, productID int, quantity int) int {

	if quantity > 10 {
		quantity = 10
	}

	productIndex := -1
	for i, product := range ps.catalogue {
		if product.ID == productID {
			productIndex = i
			break
		}
	}

	if productIndex == -1 {
		return -1
	}

	if ps.catalogue[productIndex].Availability >= quantity {
		ps.catalogue[productIndex].Availability -= quantity
		premiumProduct := 0
		if ps.catalogue[productIndex].Category == "Premium" {
			premiumProduct = 1
		}
		orderID := orderService.AddOrder(productID, ps.catalogue[productIndex].Price*float64(quantity),quantity,premiumProduct)


		return orderID
	} else {
		return -1
	}
}
