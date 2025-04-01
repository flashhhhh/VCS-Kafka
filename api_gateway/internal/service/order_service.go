package service

import (
	"encoding/json"

	"github.com/flashhhhh/pkg/kafka"
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

/*
	Order:
	- productId: quantity
	=> map[string]int
*/
func (orderService *OrderService) SendOrder(order map[string]int) error {
	// Convert map[string]int to byte array
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// Send the byte array to Kafka
	err = orderService.kafkaProducer.SendMessage(orderService.topic, orderBytes)
	if err != nil {
		return err
	}

	return nil
}