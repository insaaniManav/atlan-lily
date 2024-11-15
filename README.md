# Real-Time Metadata Streaming Service

This repository provides a real-time metadata streaming solution that demonstrates two use cases:

1. **Inbound Service**: Processes data quality events from Kafka and stores them in Elasticsearch for search and retrieval.
2. **Outbound Service**: Listens to compliance metadata events from Kafka and logs them, simulating a notification system.

## Features

- **Event-Driven Architecture**: Powered by Kafka for high-performance real-time streaming.
- **Metadata Ingestion**: Efficiently indexes data quality events into Elasticsearch for fast search capabilities.
- **Event Notifications**: Processes compliance metadata events and logs them for downstream consumption.
- **Extensibility**: Modular design for easy extension to additional use cases.

---

## Prerequisites

Before running the services, ensure you have the following installed:

1. **Docker** and **Docker Compose** for setting up Kafka and Elasticsearch.
2. **Golang** (version 1.20 or higher) for building and running the services.

---

## Getting Started

### Step 1: Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```
### Step 2: Set Up Kafka and Elasticsearch

Start Kafka and Elasticsearch using Docker Compose:

```bash
docker-compose up -d
```

This will set up:
- **Kafka**: Running on `localhost:9092`.
- **Elasticsearch**: Running on `localhost:9200`.

---

## Inbound Service

The inbound service processes data quality events from Kafka and stores them in Elasticsearch.

### Step 3: Run the Inbound Service

Navigate to the `inbound-service` directory:

```bash
docker-compose up -d

```

#### Build and Run the Service:

```bash
go mod tidy
go build -o inbound-service main.go
./inbound-service
```


The service will now listen to the `data-quality-events` topic and store processed events in Elasticsearch.

---

### Step 4: Produce a Data Quality Event

Navigate to the `inbound-service/producer` directory and produce a sample event:
```bash
cd producer
go mod tidy
go run producer.go
```

This will send a data quality event to the `data-quality-events` Kafka topic. You can verify the event was indexed in Elasticsearch by running:
```bash
curl -X GET "localhost:9200/metadata/_search" -H 'Content-Type: application/json'
```

---

## Outbound Service

The outbound service listens for compliance metadata events on Kafka and logs them.

### Step 5: Run the Outbound Service

Navigate to the `outbound-service` directory:
```bash
cd outbound-service
```

#### Build and Run the Service:
```bash
go mod tidy
go build -o outbound-service main.go
./outbound-service
```
The service will now listen to the `metadata-notifications` topic for compliance events.

---

### Step 6: Produce a Compliance Event

Navigate to the `outbound-service/producer` directory and produce a sample compliance event:
```bash
cd producer
go mod tidy
go run compliance_producer.go
```

This will send a compliance event to the `metadata-notifications` Kafka topic. The outbound service will log the event.

---

## Testing and Verification

1. **Inbound Testing**:
    - Produce a data quality event.
    - Verify it is indexed in Elasticsearch.

2. **Outbound Testing**:
    - Produce a compliance event.
    - Check the logs of the outbound service for the processed event.

---

## Cleanup

To stop all running services:
```bash
docker-compose down
```

---

## Extending the Solution

This repository demonstrates a foundational event-driven architecture that can be extended to support additional use cases, such as:

- Enriching metadata with external data sources.
- Triggering downstream workflows based on metadata events.
- Supporting multi-tenancy and fine-grained access control.
