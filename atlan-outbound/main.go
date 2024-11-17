// main.go
package main

import (
	"atlan-outbound/constants"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type ComplianceEvent struct {
	EventID   string `json:"eventId"`
	Entity    string `json:"entity"`
	Tag       string `json:"tag"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}

func main() {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{constants.KafkaURL},
		Topic:   constants.KafkaComplianceEventTopic,
		GroupID: constants.KafkaComplianceGroupTopic,
	})
	defer kafkaReader.Close()

	fmt.Println("Listening for compliance events...")
	for {
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		var event ComplianceEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}

		// Process the compliance event (for now, just print it)
		fmt.Printf("Compliance Event: %s - %s tagged as %s\n", event.EventID, event.Entity, event.Tag)
	}
}
