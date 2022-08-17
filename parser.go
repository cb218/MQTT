package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"strconv"
	"strings"
)

// Message The structs are defined from schema.json
type Message struct {
	Location	Location 	`json:"location"`
	Reading		[]Reading 	`json:"reading"`
	SensorType	string 		`json:"sensorType"`
	Time		string 		`json:"time"`
	Tower		string 		`json:"tower"`
}
type Reading struct {
	Measurement	string		`json:"measurement"`
	Name		string		`json:"name"`
	Value		string		`json:"value"`
}
type Location struct {
	Latitude	string		`json:"lat"`
	Longitude	string		`json:"long"`
}

func Parse(file string) {
	var message Message
	data, err := ioutil.ReadFile(file)
	if err != nil {fmt.Println(err)}
	err = json.Unmarshal(data, &message)
	if err != nil {fmt.Println(err)}

	tower := message.Tower
	timeToken := strings.Split(message.Time, ":")
	utime,errt := strconv.Atoi(timeToken[1])
	if errt != nil {fmt.Println(errt)}
	timestamp := time.Unix(int64(utime), 0)
	var temperature string
	for i := range message.Reading {
		if message.Reading[i].Name == "temperature" {
			temperature = message.Reading[i].Value
		}
	}
	fmt.Printf("Tower: %s, Temperature: %s, Time %s",tower,temperature,timestamp)
}

func main() {
	Parse("example1.json")
}