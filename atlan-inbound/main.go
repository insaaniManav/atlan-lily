package main

import (
	"atlan-inbound/constants"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
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
		Brokers: []string{constants.KafkaURL},
		Topic:   constants.DataQualityEventTopic,
		GroupID: constants.DataQualityGroupId,
	})
	defer kafkaReader.Close()

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

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

		fmt.Printf("Received event: %+v\n", event)

		indexName := constants.DataQualityEventIndexName
		eventJSON, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshalling event: %v", err)
			continue
		}

		req := esapi.IndexRequest{
			Index:      indexName,
			DocumentID: event.EventID, // Optional: Use the event ID as the document ID
			Body:       bytes.NewReader(eventJSON),
			Refresh:    "true", // Ensures the document is immediately searchable
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Printf("Error indexing document: %v", err)
			continue
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error response from Elasticsearch: %s", res.String())
		} else {
			log.Printf("Indexed event ID: %s", event.EventID)
		}
	}
}
