package pubsub

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
)

var pConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Publisher connected")
}

var pConnectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Publisher connection Lost: %s\n", err.Error())
}

type Publisher struct {
	broker   string
	clientID string
	Client   mqtt.Client
}

func NewPublisher(broker string, clientID string) (p Publisher) {
	p.broker = broker
	p.clientID = clientID
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientID)
	options.OnConnect = pConnectHandler
	options.OnConnectionLost = pConnectionLostHandler

	p.Client = mqtt.NewClient(options)
	if token := p.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return p
}

func (p *Publisher) PublishFile(file string, topic string) {
	// Input to be published

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	message := string(data)
	token := p.Client.Publish(topic, 0, false, message)
	token.Wait()
	fmt.Println("Input file published...")
}
