// producer.go
package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "data-quality-events",
	})
	defer writer.Close()

	event := map[string]string{
		"eventId":   "dq-1001",
		"source":    "MonteCarlo",
		"table":     "customer_data",
		"issue":     "null_values",
		"severity":  "high",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	msg, _ := json.Marshal(event)
	if err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	}); err != nil {
		log.Fatalf("Error writing message: %v", err)
	}

	log.Println("Produced data quality event to Kafka")
}
