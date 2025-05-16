package kafka

import (
	"context"
	"log"

	schema "notification-service/avro"
	"notification-service/model"

	avro "github.com/hamba/avro/v2"
	"github.com/segmentio/kafka-go"
)

type EventConsumer struct {
	reader *kafka.Reader
	schema avro.Schema
}

func NewConsumer(broker, topic, groupID string) *EventConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e2, // 1KB
		MaxBytes: 10e6, // 10MB,
	})

	schemaObj, err := avro.Parse(schema.TransferEventSchema)
	if err != nil {
		log.Fatalf("Failed to parse Avro schema: %v", err)
	}

	return &EventConsumer{reader: reader, schema: schemaObj}
}

func (c *EventConsumer) Start(ctx context.Context) {
	log.Println("🟢 Notification consumer started...")

	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		var event model.TransferEvent
		err = avro.Unmarshal(c.schema, m.Value, &event)
		if err != nil {
			log.Println("Avro unmarshal error:", err)
			continue
		}

		log.Printf("📨 NOTIFY: Transfer from %s to %s of $%.2f\n",
			event.FromAccount, event.ToAccount, event.Amount)
	}
}
