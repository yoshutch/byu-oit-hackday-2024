package events

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type EventAdapter struct {
	reader *kafka.Reader
}

func NewEventAdapter() (*EventAdapter, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "profile-updated",
		Partition: 0,
		MaxBytes:  10e6,
	})
	//err := r.SetOffset(42)
	//if err != nil {
	//	return nil, err
	//}

	return &EventAdapter{r}, nil
}

func (a EventAdapter) Listen() error {
	log.Printf("Event Adapter listening for topic: 'profile-updated' ...")
	for {
		m, err := a.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	return nil
}

func (a EventAdapter) Close() {
	log.Printf("Closing event adapter")
	if err := a.reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
