package main

import (
	"fmt"
	"headers/consumer"
	"headers/publisher"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const RabbitMqUrl = "amqp://dgkwon:test001@192.168.56.1:5672/"
const ExchangeName = "topic_test_exchange"

var consumerMsgs map[string]string
var mutex = &sync.Mutex{}

func reviceMsgHandler(name string, msg interface{}) {
	reviceMsg := msg.(amqp.Delivery)
	mutex.Lock()
	if val, ok := consumerMsgs[name]; ok {
		consumerMsgs[name] = val + ", " + reviceMsg.MessageId
	} else {
		consumerMsgs[name] = reviceMsg.MessageId
	}
	mutex.Unlock()
}

func StartConsumers() {
	fmt.Println("\n\nStart Consumers!!!!!!!!!!!!!!!!!!!!!!!")
	var wg sync.WaitGroup

	//Consumer1
	wg.Add(1)
	go func() {
		defer wg.Done()

		con1 := consumer.NewCon(
			RabbitMqUrl,
			"consumber:1",
			ExchangeName,
			"ucl.one",
			"",
			map[string]interface{}{
				"x-match": "any",
				"country": "us",
				"city":    "cd",
			})
		defer con1.Close()
		con1.Connection()
		con1.OpenChannel()
		con1.Bind(reviceMsgHandler)
	}()

	//Consumer2
	wg.Add(1)
	go func() {
		defer wg.Done()
		con2 := consumer.NewCon(
			RabbitMqUrl,
			"consumber:2",
			ExchangeName,
			"ucl.one",
			"",
			map[string]interface{}{
				"x-match": "any",
				"country": "us",
				"city":    "cd",
			})
		defer con2.Close()
		con2.Connection()
		con2.OpenChannel()
		con2.Bind(reviceMsgHandler)
	}()

	wg.Add(1)
	go func() {
		con3 := consumer.NewCon(RabbitMqUrl, "consumber:3", ExchangeName, "ucl.two", "",
			map[string]interface{}{
				"x-match": "all",
				"country": "bd",
				"city":    "cd",
			})
		defer con3.Close()
		con3.Connection()
		con3.OpenChannel()
		con3.Bind(reviceMsgHandler)
	}()
	wg.Wait()
}

func StartPublisher() {
	fmt.Println("\n\nStart Publisher!!!!!!!!!!!!!!!!!!!!!!!")

	//publisher1
	pub := publisher.NewPub(RabbitMqUrl, "publisher:1")
	defer pub.Close()
	pub.Connection()
	pub.OpenChannel()

	//Msg1
	pub.Publish(
		ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg1",
			Headers: map[string]interface{}{
				"country": "us",
				"city":    "ab",
			},
			Body: []byte(`{"username":"sysed"}`),
		},
	)

	//Msg2
	pub.Publish(
		ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg2",
			Headers: map[string]interface{}{
				"country": "us",
				"city":    "cd",
			},
			Body: []byte(`{"username":"sirajul"}`),
		},
	)

	//Msg3
	pub.Publish(
		ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg3",
			Headers: map[string]interface{}{
				"country": "uk",
				"city":    "ab",
			},
			Body: []byte(`{"username":"islam"}`),
		},
	)

	//Msg4
	pub.Publish(
		ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg4",
			Headers: map[string]interface{}{
				"country": "bd",
				"city":    "cd",
			},
			Body: []byte(`{"username":"anik", "old":"syed"}`),
		},
	)
}

func main() {
	consumerMsgs = make(map[string]string)
	go StartConsumers()
	time.Sleep(time.Second * 3)
	StartPublisher()

	fmt.Println("====== [result] ======")
	for con, msg := range consumerMsgs {
		fmt.Printf("%v: %v\n", con, msg)
	}
}
