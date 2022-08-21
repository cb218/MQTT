package pubsub

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

var sMessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Subscriber received a message on topic %s\n", msg.Topic())
	fmt.Println("Parsing message...")
	Parse(msg.Payload())
}

var sConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Subscriber connected")
}

var sConnectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Subscriber connection Lost: %s\n", err.Error())
}

type Subscriber struct {
	broker   string
	clientID string
	Client   mqtt.Client
}

func NewSubscriber(broker string, clientID string, topic string) {
	var s Subscriber
	s.broker = broker
	s.clientID = clientID
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientID)
	options.SetDefaultPublishHandler(sMessagePubHandler)
	options.OnConnect = sConnectHandler
	options.OnConnectionLost = sConnectionLostHandler

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
