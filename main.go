package main

import (
	"MQTT/pubsub"
	"time"
)

func main() {
	broker := "tcp://localhost:1883"
	topic := "test"
	pubsub.NewSubscriber(broker, "Localhost-Subscriber", topic)
	publisher := pubsub.NewPublisher(broker, "Localhost-Publisher")
	file := "example.json"
	publisher.PublishFile(file, topic)
	time.Sleep(1 * time.Second) // Gives the subscriber time to read the message
}
