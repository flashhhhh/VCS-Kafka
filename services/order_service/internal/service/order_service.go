package service

import (
	"log"
	"order_service/internal/domain"
	"order_service/internal/repository"
	"strconv"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) CreateOrder(userID int, orderInfo map[string]any) error {
	log.Printf("Creating order with info: %v", orderInfo)

	orderID, err := s.orderRepository.GetMaxID()
	if err != nil {
		log.Printf("Error getting max ID: %v", err)
	}

	orderID += 1

	var orders []domain.Order
	for id, quantity := range orderInfo {
		productID, _ := strconv.Atoi(id)

		orders = append(orders, domain.Order{
			ID:       orderID,
			UserID:   userID,
			ProductID: productID,
			Quantity:  int(quantity.(float64)),
		})
	}

	return s.orderRepository.CreateOrder(orders)
}

func (s* OrderService) CreateDetailedBill(orderInfo map[string]any) (map[string]any, error) {
	var totalPrice float64
	var orderDetails []map[string]any

	for productIDStr, quantity := range orderInfo {
		log.Printf("Calculating price for product ID: %s", productIDStr)

		productID, _ := strconv.Atoi(productIDStr)
		product, err := s.orderRepository.GetProductByID(productID)
		if err != nil {
			return nil, err
		}

		totalPrice += product.Price * float64(int(quantity.(float64)))
		orderDetails = append(orderDetails, map[string]any{
			"productID": product.ID,
			"name":      product.Name,
			"quantity":  int(quantity.(float64)),
			"price":     product.Price,
		})
	}

	return map[string]any{
		"orderDetails": orderDetails,
		"totalPrice": totalPrice,
	}, nil
}