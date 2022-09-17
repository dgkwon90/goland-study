// topic exchange publish, consumer 테스트 소스 이다.

package main

import (
	"fmt"
	"rabbitmq/consumer"
	"rabbitmq/publisher"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	// rabbitmq broker 가 설치되어 있는 테스트 서버 URL 정보
	rabbitMqUrl  = "amqp://dgkwon:test001@192.168.56.1:5672/"
	exchangeName = "topic_test_exchange"
	exchangeType = "topic"
)

// consumer가 받은 메세지를 저장 하는 map
var consumerMsgs map[string]string
var mutex = &sync.Mutex{}

// consumer가 메세지를 수신 받아 호출하는 메시지 핸들러
// 테스트에서는 메세지를 cousumer 이름을 key 수신 받은 메세지를 map에 저장한다
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

// consumer 생성
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
			"ucl",
			"user.created.*",
			nil,
		)
		defer con1.Close()
		con1.Connection()
		con1.OpenChannel()
		con1.Bind(exchangeType, reviceMsgHandler)
	}()

	// Consumer2
	wg.Add(1)
	go func() {
		defer wg.Done()
		con2 := consumer.New(
			rabbitMqUrl,
			"consumer:2",
			exchangeName,
			"uul",
			"user.updated.#",
			nil,
		)
		defer con2.Close()
		con2.Connection()
		con2.OpenChannel()
		con2.Bind(exchangeType, reviceMsgHandler)
	}()

	// Consumer3
	wg.Add(1)
	go func() {
		con3 := consumer.New(
			rabbitMqUrl,
			"consumer:3",
			exchangeName,
			"ucl.two",
			"user.#",
			nil,
		)
		defer con3.Close()
		con3.Connection()
		con3.OpenChannel()
		con3.Bind(exchangeType, reviceMsgHandler)
	}()

	// Consumer4
	wg.Add(1)
	go func() {
		con4 := consumer.New(
			rabbitMqUrl,
			"consumer:4",
			exchangeName,
			"ucl",
			"user.created.*",
			nil,
		)
		defer con4.Close()
		con4.Connection()
		con4.OpenChannel()
		con4.Bind(exchangeType, reviceMsgHandler)
	}()

	// Consumer5
	wg.Add(1)
	go func() {
		con5 := consumer.New(
			rabbitMqUrl,
			"consumer:5",
			exchangeName,
			"ua.us",
			"user.*.us",
			nil,
		)
		defer con5.Close()
		con5.Connection()
		con5.OpenChannel()
		con5.Bind(exchangeType, reviceMsgHandler)
	}()
	wg.Wait()
}

// publisher 생성 및 메세지 발신
func StartPublisher() {
	fmt.Println("============ [Start Publisher] ============")

	// publisher1
	pub := publisher.New(rabbitMqUrl, "publisher:1")
	defer pub.Close()
	pub.Connection()
	pub.OpenChannel()

	// Msg1
	pub.Publish(
		exchangeName,
		"user.created.us",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg1",
			Headers:     nil,
			Body:        []byte(`{"username":"sysed"}`),
		},
	)

	// Msg2
	pub.Publish(
		exchangeName,
		"user.created.uk",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg2",
			Headers:     nil,
			Body:        []byte(`{"username":"sirajul"}`),
		},
	)

	// Msg3
	pub.Publish(
		exchangeName,
		"user.created.bd",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg3",
			Headers:     nil,
			Body:        []byte(`{"username":"islam"}`),
		},
	)

	// Msg4
	pub.Publish(
		exchangeName,
		"user.updated.us",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg4",
			Headers:     nil,
			Body:        []byte(`{"username":"anik", "old":"syed"}`),
		},
	)

	// Msg5
	pub.Publish(
		exchangeName,
		"user.deleted.bd",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg5",
			Headers:     nil,
			Body:        []byte(`{"username":"islam"}`),
		},
	)

	// Msg6
	pub.Publish(
		exchangeName,
		"user",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			MessageId:   "Msg6",
			Headers:     nil,
			Body:        []byte(`{"username":"islam"}`),
		},
	)
}

func main() {
	fmt.Println("============ [topic exchange test] ============")

	consumerMsgs = make(map[string]string)
	go StartConsumers()
	time.Sleep(time.Second * 3)
	StartPublisher()

	fmt.Println("============ [result] ============")
	for con, msg := range consumerMsgs {
		fmt.Printf("%v: %v\n", con, msg)
	}
}
