package consumer

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Con struct {
	Url        string
	Name       string
	Exchange   string
	QueueName  string
	RoutingKey string
	Headers    map[string]interface{}
	Conn       *amqp.Connection
	Channel    *amqp.Channel
}

type ReciveMsgHandler func(name string, msg interface{})

func NewCon(url, name, exchange, queueName, routingKey string, headers map[string]interface{}) *Con {
	con := new(Con)
	con.Url = url
	con.Name = name
	con.Exchange = exchange
	con.QueueName = queueName
	con.RoutingKey = routingKey
	con.Headers = headers
	fmt.Printf("[%v] New Con :%v\n", con.Name, con)
	return con
}

func (c *Con) Connection() error {
	conn, connErr := amqp.Dial(c.Url)
	if connErr != nil {
		fmt.Printf("[%v] Rebbit MQ Connection Fail %v\n", c.Name, connErr.Error())
		return connErr
	}

	c.Conn = conn
	return nil
}

func (c *Con) OpenChannel() error {
	ch, chErr := c.Conn.Channel()
	if chErr != nil {
		fmt.Printf("[%v] Channel Fail %v\n", c.Name, chErr.Error())
		return chErr
	}
	c.Channel = ch
	return nil
}

func (c *Con) Close() {
	if c.Channel != nil {
		c.Channel.Close()
	}
	if c.Conn != nil {
		c.Conn.Close()
	}
}

func (c *Con) Bind(handler ReciveMsgHandler) error {
	exchangeDeclareErr := c.Channel.ExchangeDeclare(
		c.Exchange, // name
		"direct",   // type
		false,      // durable
		true,       // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if exchangeDeclareErr != nil {
		fmt.Printf("[%v] ExchangeDeclare Error %v\n", c.Name, exchangeDeclareErr)
		return exchangeDeclareErr
	}

	_, queueDeclareErr := c.Channel.QueueDeclare(
		c.QueueName, // name
		false,       // durable
		true,        // auto-deleted
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if queueDeclareErr != nil {
		fmt.Printf("[%v] QueueDeclare Error %v\n", c.Name, queueDeclareErr)
		return queueDeclareErr
	}

	bindErr := c.Channel.QueueBind(
		c.QueueName,  // name
		c.RoutingKey, // key(routing)
		c.Exchange,   // exchange
		false,        // no-wait
		c.Headers,    // arguments
	)
	if bindErr != nil {
		fmt.Printf("[%v] QueueBind Error %v\n", c.Name, bindErr)
		return bindErr
	}

	messages, consumeErr := c.Channel.Consume(
		c.QueueName, // queue
		c.Name,      // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if consumeErr != nil {
		fmt.Printf("[%v] Consume Error %v\n", c.Name, consumeErr)
		return consumeErr
	}

	fmt.Printf("[%v] Start wait message...\n", c.Name)
	for msg := range messages {
		handler(c.Name, msg)
		msg.Ack(true)
	}
	return nil
}
