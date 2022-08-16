package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Tower       string
	Temperature string
	Time        string
}

func main() {
	var message []Message
	Data := []byte(`
	[
		{"Tower": "Japan", "Temperature": "Tokyo", "Time": "Asia"},
	]`)

	err := json.Unmarshal(Data, &message)
	if err != nil {
		fmt.Println(err)
	}

	for i := range message {
		fmt.Println(message[i].Temperature)
	}
}
