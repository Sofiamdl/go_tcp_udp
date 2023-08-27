package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

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
	Name     string `json:"name"`
	SentTime string `json:"sentTime"`
}

func main() {
	startTime := time.Now()

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

	var executionsWaitGroup sync.WaitGroup
	executionsWaitGroup.Add(execQuant)

	// inicializa os consumidores
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

			log.Printf("🌀-%d aguardando nomes", id)

			for d := range msgs {
				var message Message
				err := json.Unmarshal(d.Body, &message)
				failOnError(err, "Failed to decode the JSON message")

				ninjaName := TransformName(message.Name)
				timestamp, err := time.Parse(time.RFC3339, message.SentTime)
				failOnError(err, "Failed to parse the message sent time")

				elapsedTime := time.Now().Sub(timestamp).Milliseconds()
				log.Printf("🥷-%d: \"olá, %s! você enviou uma mensagem %d ms atrás\"", id, ninjaName, elapsedTime)
				fmt.Println(elapsedTime)
				executionsWaitGroup.Done()
			}
		}(i)
	}
	executionsWaitGroup.Wait()

	totalElapsedTime := time.Now().Sub(startTime).Microseconds()
	log.Printf("Tempo total: %d ms", totalElapsedTime)
}

var consumersQuant int
var execQuant int

func init() {
	flag.IntVar(&consumersQuant, "consumers", 1, "number of consumers")
	flag.IntVar(&execQuant, "executions", 1, "number of executions")
	flag.Parse()
}
