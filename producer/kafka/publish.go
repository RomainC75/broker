package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func NewConnection() {
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	kafkaUrl := fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)
	fmt.Println("kafka url : ", kafkaUrl)

	topic := os.Getenv("KAFKA_TOPIC")

	var err error
	l := log.New(os.Stdout, "kafka writer: ", 0)
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaUrl},
		Topic:   topic,
		Logger:  l,
	})
	if err != nil {
		log.Fatal("KAFKA Producer // failed to dial leader:", err)
	}
}

func Produce(ctx context.Context, i int, message string) {
	// to produce messages
	if writer == nil {
		fmt.Println("no wriiter")
	}
	err := writer.WriteMessages(ctx, kafka.Message{
		Partition: 0,
		Key:       []byte(strconv.Itoa(i)),
		Value:     []byte("this is message : " + message),
	})
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
