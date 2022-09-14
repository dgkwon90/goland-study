package consumer_test

import (
	"fanout/consumer"
	"fmt"
	"sync"
	"testing"
)

//var reviceMsgs map[string]string

// func handler(name string, msg map[string]interface{}) {
// 	reviceMsgs[name] = reviceMsgs[name] + "," + msg["msg"].(string)
// 	for key, value := range reviceMsgs {
// 		fmt.Println(key, " : ", value)
// 	}
// }

func TestConsumber(t *testing.T) {
	rabbitMqUrl := "amqp://dgkwon:test001@192.168.56.1:5672/"
	exchangeName := "fanout_test_exchange"
	//reviceMsgs = make(map[string]string)

	fmt.Println("Start Consumers!!!!!!!!!!!!!!!!!!!!!!!")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		con1 := consumer.NewCon(rabbitMqUrl, "consumber:1", exchangeName, "ucl")
		defer con1.Close()
		con1.Connection()
		con1.OpenChannel()
		con1.Bind( /*handler*/ )
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		con2 := consumer.NewCon(rabbitMqUrl, "consumber:2", exchangeName, "ucl")
		defer con2.Close()
		con2.Connection()
		con2.OpenChannel()
		con2.Bind( /*handler*/ )
	}()

	wg.Add(1)
	go func() {
		con3 := consumer.NewCon(rabbitMqUrl, "consumber:3", exchangeName, "ucl.two")
		defer con3.Close()
		con3.Connection()
		con3.OpenChannel()
		con3.Bind( /*handler*/ )
	}()
	wg.Wait()
}
