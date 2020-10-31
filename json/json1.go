package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Broker          string `json:"broker"`
	ResultBackend   string `json:"result_backend"`
	ResultsExpireIn int    `json:"results_expire_in"`
	Exchange        string `json:"exchange"`
	ExchangeType    string `json:"exchange_type"`
	DefaultQueue    string `json:"default_queue"`
	BindingKey      string `json:"binding_key"`
}

func main() {
	var cnf = Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		ResultBackend: "amqp",
		Exchange:      "machinery_exchange",
		ExchangeType:  "direct",
		DefaultQueue:  "machinery_tasks",
		BindingKey:    "machinery_task",
	}

	body, err := json.Marshal(cnf)
	if err != nil {
		panic(err.Error())
	}

	os.Stdout.Write(body)
}
