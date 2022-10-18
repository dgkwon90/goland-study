package main

import (
	"encoding/json"
	"fmt"
	"rabbitmq/consumer"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitMqUrl  = "amqp://dgkwon:test001@192.168.56.1:5672/"
	exchangeName = "device-connection-topic"
	exchangeType = "fanout"
)

func receiveMsgHandler(name string, msg interface{}) {
	reviceMsg := msg.(amqp.Delivery)
	jsonBody := new(map[string]interface{})
	err := json.Unmarshal(reviceMsg.Body, jsonBody)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Printf("[%v] receive json: %v\n", name, jsonBody)
}

func StartConsumers() {
	fmt.Println("============ [Start Consumers] ============")
	var wg sync.WaitGroup

	// Consumer1
	wg.Add(1)
	go func() {
		defer wg.Done()
		con1 := consumer.New(
			rabbitMqUrl,
			"consumer:1",
			exchangeName,
			exchangeName+"-pull",
			"",
			nil,
		)
		defer con1.Close()
		con1.Connection()
		con1.OpenChannel()
		con1.Bind(exchangeType, receiveMsgHandler)
	}()

	// // Consumer2
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	con2 := consumer.New(
	// 		rabbitMqUrl,
	// 		"consumer:2",
	// 		exchangeName,
	// 		exchangeName+"-pull",
	// 		"",
	// 		nil,
	// 	)
	// 	defer con2.Close()
	// 	con2.Connection()
	// 	con2.OpenChannel()
	// 	con2.Bind(exchangeType, receiveMsgHandler)
	// }()

	// // Consumer3
	// wg.Add(1)
	// go func() {
	// 	con3 := consumer.New(
	// 		rabbitMqUrl,
	// 		"consumer:3",
	// 		exchangeName,
	// 		exchangeName+"-pull",
	// 		"",
	// 		nil,
	// 	)
	// 	defer con3.Close()
	// 	con3.Connection()
	// 	con3.OpenChannel()
	// 	con3.Bind(exchangeType, receiveMsgHandler)
	// }()

	// Consumer4
	// wg.Add(1)
	// go func() {
	// 	con4 := consumer.New(
	// 		rabbitMqUrl,
	// 		"consumer:4",
	// 		exchangeName,
	// 		exchangeName+"-pull",
	// 		"",
	// 		nil,
	// 	)
	// 	defer con4.Close()
	// 	con4.Connection()
	// 	con4.OpenChannel()
	// 	con4.Bind(exchangeType, receiveMsgHandler)
	// }()
	wg.Wait()
}

func main() {
	StartConsumers()
}
