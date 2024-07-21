package kafka

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/websocket"
)

var writer *kafka.Writer

var connection *Connection

type Connection struct {
	url    url.URL
	config *websocket.Config
	conn   *websocket.Conn
}

var (
	u = url.URL{Scheme: "ws", Host: "localhost:3005", Path: "/ws"}
)

// func NewConnection() {
// 	kafkaHost := os.Getenv("KAFKA_HOST")
// 	kafkaPort := os.Getenv("KAFKA_PORT")
// 	kafkaUrl := fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)
// 	fmt.Println("kafka url : ", kafkaUrl)

// 	topic := os.Getenv("KAFKA_TOPIC")

// 	var err error
// 	l := log.New(os.Stdout, "kafka writer: ", 0)
// 	writer = kafka.NewWriter(kafka.WriterConfig{
// 		Brokers: []string{kafkaUrl},
// 		Topic:   topic,
// 		Logger:  l,
// 	})
// 	if err != nil {
// 		log.Fatal("KAFKA Producer // failed to dial leader:", err)
// 	}
// }

func NewConn() *Connection {

	config, err := websocket.NewConfig(u.String(), "http://localhost")
	if err != nil {
		log.Fatal("error with config: ", err.Error())
	}
	ctx := context.Background()
	conn, err := config.DialContext(ctx)
	if err != nil {
		log.Fatal("error trying to dial: ", err.Error())
	}

	connection = &Connection{
		url:    u,
		config: config,
		conn:   conn,
	}
	return connection
}

func Produce(ctx context.Context, i int, message string) {
	// to produce messages
	if connection == nil {
		fmt.Println("no wriiter")
	}
	_, err := connection.conn.Write([]byte(message))
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
