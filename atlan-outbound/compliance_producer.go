// compliance_producer.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "metadata-notifications",
	})
	defer writer.Close()

	event := map[string]string{
		"eventId":   "compliance-2001",
		"entity":    "customer_data",
		"tag":       "PII",
		"action":    "update",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	msg, _ := json.Marshal(event)
	if err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	}); err != nil {
		log.Fatalf("Error writing message: %v", err)
	}

	log.Println("Produced compliance event to Kafka")
}
