// fanout exchange publish, consumer 테스트 소스 이다.

package main

import (
	"encoding/json"
	"fmt"
	"rabbitmq/consumer"
	"rabbitmq/publisher"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitMqUrl         = "amqp://dgkwon:test001@192.168.56.1:5672/"
	receiveExchangeName = "connection-request-topic"
	exchangeType        = "fanout"
	publishExchangeName = "xmpp-server.test"
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
			receiveExchangeName,
			receiveExchangeName+"-pull",
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
	// 		nil)
	// 	defer con2.Close()
	// 	con2.Connection()
	// 	con2.OpenChannel()
	// 	con2.Bind(exchangeType, receiveMsgHandler)
	// }()

	// // Consumer2
	// wg.Add(1)
	// go func() {
	// 	con3 := consumer.New(
	// 		rabbitMqUrl,
	// 		"consumer:3",
	// 		exchangeName,
	// 		exchangeName+"-pull",
	// 		"",
	// 		nil)
	// 	defer con3.Close()
	// 	con3.Connection()
	// 	con3.OpenChannel()
	// 	con3.Bind(exchangeType, receiveMsgHandler)
	// }()
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
		publishExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "Application/json",
			//MessageId:   "Msg1",
			Headers: map[string]interface{}{
				"method":     "connreq",
				"endPointID": "123456-stb-00001",
				"taskId":     "test001",
				"topicId":    receiveExchangeName,
			},
			// Body: []byte(`
			// {
			// 	"method":"connreq",
			// 	"endPointID":"connreq",
			// 	"taskId":"test001",
			// 	"topicId":"xmpp-server.2e7b1ea277d5421ea8affaf9d29ee6db",
			// }
			// `),
		},
	)
}

func main() {
	go StartConsumers()
	time.Sleep(time.Second * 3)
	StartPublisher()
	time.Sleep(time.Second * 30)
}
