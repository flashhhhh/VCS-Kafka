package main

import (
	"encoding/json"
	"log"
	"order_service/infrastructure/postgres"
	"order_service/internal/repository"
	"order_service/internal/service"
	"os"
	"strconv"

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
	envFile := "config/order." + environment + ".env"
	env.LoadEnv(envFile)
	log.Println("Environment variables loaded from file:", envFile)

	// Database connection
	dsn := "host=" + env.GetEnv("ORDER_DB_HOST", "localhost") +
		" user=" + env.GetEnv("ORDER_DB_USER", "root") +
		" password=" + env.GetEnv("ORDER_DB_PASSWORD", "password") +
		" dbname=" + env.GetEnv("ORDER_DB_NAME", "order_db") +
		" port=" + env.GetEnv("ORDER_DB_PORT", "5432") +
		" sslmode=disable"
	
	db, err := postgres.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = postgres.MigrateDB(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Kafka
	brokers := []string{env.GetEnv("KAFKA_HOST", "localhost") + ":" + env.GetEnv("KAFKA_PORT", "9092")}
	log.Printf("Kafka brokers: %v", brokers)
	
	topic := "order_topic"
	kafkaConsumer, err := kafka.NewKafkaConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	kafkaProducer, err := kafka.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	// Initialize repositories
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(*orderRepository)

	// Consume messages from the topic
	for {
		partitionConsumer, err := kafkaConsumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Failed to consume partition: %v", err)
		}

		for message := range partitionConsumer.Messages() {
			log.Printf("Received message: %s", string(message.Value))
			
			// Process the message
			// Convert the message to map[string]any
			var orderData map[string]any
			err = json.Unmarshal(message.Value, &orderData)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			orderIDStr := orderData["userID"].(string)
			orderID, _ := strconv.Atoi(orderIDStr)

			// Bill
			orderDetails, err := orderService.CreateDetailedBill(orderData["order"].(map[string]any))
			if err != nil {
				log.Printf("Failed to create detailed bill: %v", err)
				continue
			}
			log.Printf("Detailed bill created successfully: %v", orderDetails)

			// Insert order into the database
			creationErr := orderService.CreateOrder(orderID, orderData["order"].(map[string]any))
			if creationErr != nil {
				log.Printf("Failed to create order: %v", creationErr)
				continue
			}

			log.Printf("Order created successfully for user ID: %d", orderID)

			// Send the order details to api_gateway_topic
			orderDetails["orderID"] = orderData["orderID"]

			orderDetailsJSON, err := json.Marshal(orderDetails)
			if err != nil {
				log.Printf("Failed to marshal order details: %v", err)
				continue
			}

			err = kafkaProducer.SendMessage("api_gateway_topic", orderDetailsJSON)
			if err != nil {
				log.Printf("Failed to send message to api_gateway_topic: %v", err)
				continue
			}
			log.Printf("Order details sent to api_gateway_topic successfully")
		}
	}
}