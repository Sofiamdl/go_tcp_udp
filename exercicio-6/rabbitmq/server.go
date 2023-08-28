package main

import (
	"encoding/json"
	"log"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ninjaDict = map[string]string{
	"a": "ka",
	"b": "zu",
	"c": "mi",
	"d": "te",
	"e": "ku",
	"f": "lu",
	"g": "ji",
	"h": "ri",
	"i": "ki",
	"j": "zu",
	"k": "me",
	"l": "ta",
	"m": "rin",
	"n": "to",
	"o": "mo",
	"p": "no",
	"q": "ke",
	"r": "shi",
	"s": "ari",
	"t": "chi",
	"u": "do",
	"v": "ru",
	"w": "mei",
	"x": "na",
	"y": "fu",
	"z": "zi",
}

func TransformName(name string) string {
	newName := ""
	for i := 0; i < len(name); i++ {
		letter := strings.ToLower(string(name[i]))
		if letter == " " {
			newName += " "
			continue
		}
		newName += ninjaDict[letter]
	}
	return newName
}

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

	requestQueue, err := ch.QueueDeclare(
		"ninjaNameRequest", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")

	// log.Printf("🥷servidor aguardando nomes🥷")

	for d := range msgs {
		var request Message
		err := json.Unmarshal(d.Body, &request)
		failOnError(err, "Failed to decode the JSON message")

		ninjaName := TransformName(request.Name)
		reply := Message{
			Name: ninjaName,
		}
		replyJson, err := json.Marshal(reply)
		failOnError(err, "Failed to parse the JSON message")

		err = ch.Publish(
			"",
			d.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: d.CorrelationId, // usa correlation id do request
				Body:          replyJson,
			},
		)
		failOnError(err, "Failed to send the JSON reply message")
	}
}
