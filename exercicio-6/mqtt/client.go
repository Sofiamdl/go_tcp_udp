package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Message struct {
	Name string    `json:"name"`
	Id   int       `json:"id"`
	Time time.Time `json:"time"`
}

var startTime time.Time

var timesArray [200]time.Time

const qosClient = 1

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("mqtt://localhost:1883")
	opts.SetClientID("cliente")

	client := MQTT.NewClient(opts)

	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	defer client.Disconnect(250)

	token = client.Subscribe("reply", qosClient, receiveHandlerClient)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for i := 0; i < 10000; i++ {
		timesArray[i] = time.Now()

		msg, err := json.Marshal(Message{Name: "sofia", Id: i, Time: time.Now()})
		failOnError(err, "Failed to parse the JSON message")

		token := client.Publish("request", qosClient, false, msg)
		token.Wait()
		if token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
		// fmt.Printf("Mensagem Publicada: %s\n", msg)
		time.Sleep(time.Microsecond * time.Duration(50000))
	}
}

var receiveHandlerClient MQTT.MessageHandler = func(c MQTT.Client, m MQTT.Message) {
	var response Message

	err := json.Unmarshal(m.Payload(), &response)
	failOnError(err, "Failed to decode the JSON message")

	endTime := time.Now().Sub(response.Time).Nanoseconds()
	fmt.Println(endTime)
	// log.Printf("olá! seu nome ninja é 🌀%s🌀", response.Name)
}
