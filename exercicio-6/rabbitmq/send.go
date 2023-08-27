package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exec := 0

	for exec < execQuant {
		select {
		case <-ctx.Done():
			return
		default:
			now := time.Now()
			data := map[string]interface{}{
				"name":     "chico fumaca",
				"sentTime": now.Format(time.RFC3339), // format timestamp as string
			}

			jsonData, err := json.Marshal(data)
			failOnError(err, "Failed to marshal JSON")

			err = ch.PublishWithContext(ctx,
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "application/json", // Use JSON content type
					Body:        jsonData,
				})
			failOnError(err, "Failed to publish a message")
			log.Printf("📤 nome \"%s\" enviado - \"%s\"!\n", data["name"], data["sentTime"])

			exec++
		}
	}
}

var execQuant int

func init() {
	flag.IntVar(&execQuant, "executions", 10000, "number of executions")
	flag.Parse()
}
