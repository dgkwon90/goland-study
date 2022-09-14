package main

import (
	"fanout/consumer"
	"fanout/publisher"
	"fmt"
	"sync"
	"time"
)

const rabbitMqUrl = "amqp://dgkwon:test001@192.168.56.1:5672/"
const exchangeName = "fanout_test_exchange"

func StartConsumers( /*doneStart chan bool*/ ) {
	fmt.Println("\n\nStart Consumers!!!!!!!!!!!!!!!!!!!!!!!")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		con1 := consumer.NewCon(rabbitMqUrl, "consumber:1", exchangeName, "ucl")
		defer con1.Close()
		con1.Connection()
		con1.OpenChannel()
		con1.Bind()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		con2 := consumer.NewCon(rabbitMqUrl, "consumber:2", exchangeName, "ucl")
		defer con2.Close()
		con2.Connection()
		con2.OpenChannel()
		con2.Bind()
	}()

	wg.Add(1)
	go func() {
		con3 := consumer.NewCon(rabbitMqUrl, "consumber:3", exchangeName, "ucl.two")
		defer con3.Close()
		con3.Connection()
		con3.OpenChannel()
		con3.Bind()
	}()
	wg.Wait()
}

func StartPublisher() {
	fmt.Println("\n\nStart Publisher!!!!!!!!!!!!!!!!!!!!!!!")
	pub := publisher.NewPub(rabbitMqUrl, "publisher:1", exchangeName)
	defer pub.Close()
	pub.Connection()
	pub.OpenChannel()
	pub.Publish(
		map[string]interface{}{
			"Msg": 1,
		},
		[]byte("{\"username\":\"sysed\"}"),
	)
	time.Sleep(time.Second * 1)
	pub.Publish(
		map[string]interface{}{
			"Msg": 2,
		},
		[]byte("{\"username\":\"sirajul\"}"),
	)
	time.Sleep(time.Second * 1)
	pub.Publish(
		map[string]interface{}{
			"Msg": 3,
		},
		[]byte("{\"username\":\"islam\"}"),
	)
	time.Sleep(time.Second * 1)
	pub.Publish(
		map[string]interface{}{
			"Msg": 4,
		},
		[]byte("{\"username\":\"anik\", \"old\":\"syed\"}"),
	)
	time.Sleep(time.Second * 1)
	pub.Publish(
		map[string]interface{}{
			"Msg": 5,
		},
		[]byte("{\"username\":\"ssi-anik\"}"),
	)
}

func main() {
	go StartConsumers()
	time.Sleep(time.Second * 3)
	StartPublisher()
}
