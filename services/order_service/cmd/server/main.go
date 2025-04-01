package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/flashhhhh/pkg/env"
	"github.com/flashhhhh/pkg/kafka"
)

func main() {
	var environment string

	args := os.Args
	if len(args) > 1 {
		environment = args[1]
		println("Running in environment: ", environment)
	} else {
		println("No environment specified, defaulting to local")
		environment = "local"
	}

	// Load environment variables from the specified file
	envFile := "config/" + environment + ".env"
	env.LoadEnv(envFile)
	log.Println("Environment variables loaded from file:", envFile)

	// Kafka
	brokers := []string{env.GetEnv("KAFKA_HOST", "localhost") + ":" + env.GetEnv("KAFKA_PORT", "9092")}
	topic := "order_topic"
	kafkaConsumer, err := kafka.NewKafkaConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// Consume messages from the topic
	for {
		partitionConsumer, err := kafkaConsumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Failed to consume partition: %v", err)
		}

		for message := range partitionConsumer.Messages() {
			log.Printf("Received message: %s", string(message.Value))
			// Process the message
			// ...
		}
	}
}