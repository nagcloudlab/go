package kafka

import (
	"context"
	"log"
	"time"

	schema "transfer-service/avro" // your local Avro schema
	"transfer-service/model"       // TransferEvent struct

	avro "github.com/hamba/avro/v2" // Avro library
	"github.com/segmentio/kafka-go" // Kafka producer
)

type EventProducer struct {
	writer *kafka.Writer
	schema avro.Schema
}

// NewProducer creates a Kafka writer and parses the Avro schema
func NewProducer(broker, topic string) *EventProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	// Parse the Avro schema string
	schemaObj, err := avro.Parse(schema.TransferEventSchema)
	if err != nil {
		log.Fatalf("Failed to parse Avro schema: %v", err)
	}

	return &EventProducer{
		writer: writer,
		schema: schemaObj,
	}
}

// SendTransferEvent marshals the TransferEvent and writes it to Kafka
func (p *EventProducer) SendTransferEvent(event model.TransferEvent) error {
	payload, err := avro.Marshal(p.schema, event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.FromAccount),
		Value: payload,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(context.Background(), msg)
}
