package pubsub

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Message The structs are defined from schema.json
type Message struct {
	Location    Location      `json:"location"`
	Measurement []Measurement `json:"measurement"`
	SensorType  string        `json:"sensorType"`
	Time        string        `json:"time"`
	Tower       string        `json:"tower"`
}
type Measurement struct {
	Measurement string `json:"measurement"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}
type Location struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"long"`
}

func Parse(data []byte) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		fmt.Println(err)
	}

	tower := message.Tower
	timeToken := strings.Split(message.Time, ":")
	utime, errt := strconv.Atoi(timeToken[1])
	if errt != nil {
		fmt.Println(errt)
	}
	timestamp := time.Unix(int64(utime), 0)
	var temperature string
	for i := range message.Measurement {
		if message.Measurement[i].Name == "temperature" {
			temperature = message.Measurement[i].Value
		}
	}
	fmt.Printf("Tower: %s, Temperature: %s, Time %s", tower, temperature, timestamp)
}
