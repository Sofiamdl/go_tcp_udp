package main

import (
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

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"namesToTransform", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// inicializa os consumidores
	consumersQuant := 3
	for i := 0; i < consumersQuant; i++ {
		go func(id int) {
			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				true,   // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			failOnError(err, "Failed to register a consumer")

			log.Printf("🌀-%d aguardando nomes. para sair aperte CTRL+C", id)

			for d := range msgs {
				ninjaName := TransformName(string(d.Body))
				log.Printf("🥷-%d: \"olá, %s!\"", id, ninjaName)
			}
		}(i)
	}

	var forever chan struct{}
	<-forever
}
