package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var dict = map[string]string{
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
		if i == 3 {
			newName += " "
		}
		newName += dict[letter]
	}
	return newName
}

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

const qos = 1

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("mqtt://localhost:1883")
	opts.SetClientID("subscriber 1")
	// opts.DefaultPublishHandler = receiveHandler

	client := MQTT.NewClient(opts)

	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	defer client.Disconnect(250)

	token = client.Subscribe("request", qos, receiveHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// log.Printf("🥷servidor aguardando nomes🥷")
	fmt.Scanln()
}

var receiveHandler MQTT.MessageHandler = func(c MQTT.Client, m MQTT.Message) {
	var request Message
	err := json.Unmarshal(m.Payload(), &request)
	failOnError(err, "Failed to decode the JSON message")

	ninjaName := TransformName(request.Name)
	response := Message{
		Name: ninjaName,
		Id:   request.Id,
		Time: request.Time,
	}
	responseJson, err := json.Marshal(response)
	failOnError(err, "Failed to parse the JSON message")

	token := c.Publish("reply", qos, false, responseJson)
	token.Wait()
	failOnError(err, "Failed to send the JSON reply message")
}
