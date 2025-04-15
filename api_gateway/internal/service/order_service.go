package service

import (
	"encoding/json"

	"github.com/flashhhhh/pkg/kafka"
	"github.com/google/uuid"
)

type OrderService struct {
	kafkaProducer kafka.KafkaProducer;
	topic string
}

func NewOrderService(kafkaProducer kafka.KafkaProducer, topic string) *OrderService {
	return &OrderService{
		kafkaProducer: kafkaProducer,
		topic:         topic,
	}
}

var orderInformation = make(map[string]map[string]any)

func (orderService *OrderService) SendOrder(orderInfo map[string]any) (string, error) {
	// Create UUID for the order
	orderID, _ := uuid.NewUUID()
	orderInfo["orderID"] = orderID.String()

	// Convert map[string]int to byte array
	orderBytes, err := json.Marshal(orderInfo)
	if err != nil {
		return "", err
	}

	// Send the byte array to Kafka
	err = orderService.kafkaProducer.SendMessage(orderService.topic, orderBytes)
	if err != nil {
		return "", err
	}

	return orderID.String(), nil
}

func (orderService *OrderService) UpdateOrder(orderInfo map[string]any) error {
	orderID := orderInfo["orderID"].(string)
	order := make(map[string]any)

	for key, value := range orderInfo {
		if key != "orderID" {
			order[key] = value
		}
	}

	orderInformation[orderID] = order
	return nil
}

func (orderService *OrderService) GetOrder(orderID string) (map[string]any) {
	order, exists := orderInformation[orderID]
	
	if !exists {
		return nil
	}

	return order
}