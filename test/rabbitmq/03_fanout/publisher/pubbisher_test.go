package publisher_test

import (
	"fanout/publisher"
	"fmt"
	"testing"
)

func TestPublisher(t *testing.T) {
	rabbitMqUrl := "amqp://dgkwon:test001@192.168.56.1:5672/"
	exchangeName := "fanout_test_exchange"

	fmt.Println("Start Publisher!!!!!!!!!!!!!!!!!!!!!!!")
	pub := publisher.NewPub(rabbitMqUrl, "publisher:1", exchangeName)
	defer pub.Close()
	pub.Connection()
	pub.OpenChannel()
	pub.Publish(
		map[string]interface{}{
			"Msg": "1",
		},
		[]byte("{\"username\":\"sysed\"}"),
	)
	pub.Publish(
		map[string]interface{}{
			"Msg": "2",
		},
		[]byte("{\"username\":\"sirajul\"}"),
	)
	pub.Publish(
		map[string]interface{}{
			"Msg": "3",
		},
		[]byte("{\"username\":\"islam\"}"),
	)
	pub.Publish(
		map[string]interface{}{
			"Msg": "4",
		},
		[]byte("{\"username\":\"anik\", \"old\":\"syed\"}"),
	)
	pub.Publish(
		map[string]interface{}{
			"Msg": "5",
		},
		[]byte("{\"username\":\"ssi-anik\"}"),
	)
}
