package events

// TODO couldn't figure out how to get confluence's kafka client to work, so using segmentio sdk
import (
	"byu.edu/hackday-profile/dto"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type EventAdapter struct {
	//producer *kafka.Producer
	writer *kafka.Writer
}

func NewEventAdapter() (*EventAdapter, error) {
	//p, err := kafka.NewProducer(&kafka.ConfigMap{
	//	"bootstrap.servers": "eventbus:9092",
	//	"client.id":         "local123",
	//})
	//if err != nil {
	//	return nil, err
	//}
	//return &EventAdapter{producer: p}, nil
	w := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		//Topic:    "profile-updated",
		Balancer: &kafka.LeastBytes{},
	}
	return &EventAdapter{w}, nil
}

func (a EventAdapter) SendProfileUpdatedEvent(profile dto.Profile) error {
	topic := "profile-updated"
	jsonObj, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	err = a.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Key:   []byte("Key-A"),
		Value: jsonObj,
	})
	if err != nil {
		return err
	}
	log.Printf("Event message sent!")
	return nil
	//a.producer.Produce(&kafka.Message{
	//	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//	Key:            make([]byte, 0),
	//	Value:          make([]byte, 0),
	//})
}
