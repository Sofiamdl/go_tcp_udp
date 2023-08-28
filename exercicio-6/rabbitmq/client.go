package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Message struct {
	Name string `json:"name"`
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	responseQueue, err := ch.QueueDeclare(
		"ninjaNameResponse", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		responseQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to register a consumer")

	for i := 0; i < 10000; i++ {
		startTime := time.Now()
		request := Message{
			Name: "sofia",
		}
		requestJson, err := json.Marshal(request)
		failOnError(err, "Failed to parse the JSON message")

		correlationID := uuid.New().String()

		err = ch.Publish(
			"",
			"ninjaNameRequest",
			false,
			false,
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: correlationID,
				ReplyTo:       responseQueue.Name,
				Body:          requestJson,
			},
		)

		m := <-msgs
		var reply Message
		err = json.Unmarshal(m.Body, &reply)
		failOnError(err, "Failed to decode the JSON message")

		// log.Printf("olá %s! seu nome ninja é 🌀%s🌀", request.Name, reply.Name)
		elapsedTime := time.Now().Sub(startTime).Milliseconds()
		fmt.Println(elapsedTime)
	}
}

var clientsQuant int

func init() {
	flag.IntVar(&clientsQuant, "clients", 1, "number of clients")
	flag.Parse()
}
