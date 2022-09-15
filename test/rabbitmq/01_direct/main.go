package main

import (
	"direct/consumer"
	"direct/publisher"
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const RabbitMqUrl = "amqp://dgkwon:test001@192.168.56.1:5672/"
const ExchangeName = "direct_test_exchange"

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
		con1 := consumer.NewCon(RabbitMqUrl, "consumber:1", ExchangeName, "ucl", "user.created", nil)
		defer con1.Close()
		con1.Connection()
		con1.OpenChannel()
		con1.Bind(reviceMsgHandler)
	}()

	//Consumer2
	wg.Add(1)
	go func() {
		defer wg.Done()
		con2 := consumer.NewCon(RabbitMqUrl, "consumber:2", ExchangeName, "uul", "user.updated", nil)
		defer con2.Close()
		con2.Connection()
		con2.OpenChannel()
		con2.Bind(reviceMsgHandler)
	}()

	//Consumer3
	wg.Add(1)
	go func() {
		con3 := consumer.NewCon(RabbitMqUrl, "consumber:3", ExchangeName, "ucl.two", "user.created", nil)
		defer con3.Close()
		con3.Connection()
		con3.OpenChannel()
		con3.Bind(reviceMsgHandler)
	}()

	//Consumer4
	wg.Add(1)
	go func() {
		con4 := consumer.NewCon(RabbitMqUrl, "consumber:4", ExchangeName, "ucl", "user.created", nil)
		defer con4.Close()
		con4.Connection()
		con4.OpenChannel()
		con4.Bind(reviceMsgHandler)
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
		"user.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg1",
			Headers:     nil,
			Body:        []byte("{\"username\":\"sysed\"}"),
		},
	)

	//Msg2
	pub.Publish(
		ExchangeName,
		"user.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg2",
			Headers:     nil,
			Body:        []byte("{\"username\":\"sirajul\"}"),
		},
	)

	//Msg3
	pub.Publish(
		ExchangeName,
		"user.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg3",
			Headers:     nil,
			Body:        []byte("{\"username\":\"islam\"}"),
		},
	)

	//Msg4
	pub.Publish(
		ExchangeName,
		"user.updated",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg4",
			Headers:     nil,
			Body:        []byte("{\"username\":\"anik\", \"old\":\"syed\"}"),
		},
	)

	//Msg5
	pub.Publish(
		"",
		"ucl.two",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg5",
			Headers:     nil,
			Body:        []byte("{\"username\":\"islam\"}"),
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
