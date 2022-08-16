package main

import (
	"context"
	"log"
	"net/url"
	"os"
	//"github.com/eclipse/paho.mqtt.golang"
)

func (mqtt *MQTT5) connect() {
	parsedURL, e := url.Parse(mqtt.serverURL)
	if e != nil {
		log.Fatal("MQTT URL parse failed: ", e)
		return
	}

	cliCfg := autopaho.ClientConfig{
		BrokerUrls:        []*url.URL{parsedURL},
		KeepAlive:         30,
		ConnectRetryDelay: 10000,
		OnConnectionUp:    func(*autopaho.ConnectionManager, *paho.Connack) { log.Info("mqtt connection up") },
		OnConnectError:    func(err error) { log.Error("error whilst attempting connection: ", err) },
		Debug:             paho.NOOPLogger{},
		ClientConfig: paho.ClientConfig{
			ClientID:      "qttBroker",
			OnClientError: func(err error) { log.Error("server requested disconnect: ", err) },
			OnServerDisconnect: func(d *paho.Disconnect) {
				if d.Properties != nil {
					log.Warn("server requested disconnect: ", d.Properties.ReasonString)
				} else {
					log.Warn("server requested disconnect; reason code: ", d.ReasonCode)
				}
			},
		},
	}
	mqtt.ctx, mqtt.cancel = context.WithCancel(context.Background())
	var err error
	mqtt.cm, err = autopaho.NewConnection(mqtt.ctx, cliCfg)
	if err != nil {
		log.Error("Connection failed ", err)
		os.Exit(-1)
	}
}

func main() {
	connect()
}
