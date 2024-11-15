// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type DataQualityEvent struct {
	EventID   string `json:"eventId"`
	Source    string `json:"source"`
	Table     string `json:"table"`
	Issue     string `json:"issue"`
	Severity  string `json:"severity"`
	Timestamp string `json:"timestamp"`
}

func main() {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "data-quality-events",
		GroupID: "inbound-group",
	})
	defer kafkaReader.Close()

	fmt.Println("Listening for data quality events...")
	for {
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		var event DataQualityEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}

		fmt.Println(event)
	}
}
