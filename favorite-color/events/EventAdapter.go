package events

import (
	"byu.edu/hackday-favorite-color/services"
	"byu.edu/hackday-profile/dto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type EventAdapter struct {
	reader          *kafka.Reader
	favColorService *services.FavColorService
}

func NewEventAdapter(service *services.FavColorService) (*EventAdapter, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "profile-updated",
		Partition: 0,
		MaxBytes:  10e6,
	})
	err := r.SetOffset(7)
	if err != nil {
		return nil, err
	}

	return &EventAdapter{r, service}, nil
}

func (a EventAdapter) Listen() error {
	log.Printf("Event Adapter listening for topic: 'profile-updated' ...")
	for {
		m, err := a.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		var profile dto.Profile
		err = json.Unmarshal(m.Value, &profile)
		if err != nil {
			log.Printf("Error JSON unmarshal: %s", err)
			return err
		}
		err = a.favColorService.UpdateName(1, profile)
		if err != nil {
			log.Printf("Error updating fav color name: %s", err)
			return err
		}
	}

	return nil
}

func (a EventAdapter) Close() {
	log.Printf("Closing event adapter")
	if err := a.reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
